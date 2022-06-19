package main

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"go.uber.org/zap"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ssm"

	ghinstallation "github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v45/github"

	"github.com/yunomu/checks-codepipeline/lambda/checks/db"
	"github.com/yunomu/checks-codepipeline/lambda/checks/pipelinestat"
	"github.com/yunomu/checks-codepipeline/lambda/checks/service"
)

var logger *zap.Logger

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = l
}

type appLogger struct{}

func (*appLogger) Request(event *events.CloudWatchEvent) {
	logger.Info("Request", zap.Any("request", event))
}

func (*appLogger) Error(method string, err error, execId string, checkRunId int64, state string) {
	logger.Error(method,
		zap.Error(err),
		zap.String("pipelineExecutionId", execId),
		zap.Int64("checkRunId", checkRunId),
		zap.String("state", state),
	)
}

func getEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		logger.Fatal("ENV not found", zap.String("key", key))
	}

	return v
}

func getEnvInt(key string) int64 {
	s := getEnv(key)

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		logger.Fatal("ENV parse int error",
			zap.String("key", key),
			zap.String("val", s),
		)
	}

	return i
}

func main() {
	region := getEnv("REGION")
	appID := getEnvInt("GITHUB_APP_ID")
	installationID := getEnvInt("GITHUB_INSTALLATION_ID")
	pkSecretName := getEnv("GITHUB_PK_SECRET_NAME")
	table := getEnv("PIPELINE_DB_TABLE")

	sess, err := session.NewSession()
	if err != nil {
		logger.Fatal("session.NewSession", zap.Error(err))
	}

	cfg := aws.NewConfig().WithRegion(region)

	ssmClient := ssm.New(sess, cfg)
	ctx := context.Background()

	out, err := ssmClient.GetParameterWithContext(ctx, &ssm.GetParameterInput{
		Name:           aws.String(pkSecretName),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		logger.Fatal("ssm.GetParameter", zap.Error(err))
	}
	pk := []byte(aws.StringValue(out.Parameter.Value))

	tr, err := ghinstallation.New(http.DefaultTransport, appID, installationID, pk)
	if err != nil {
		logger.Fatal("ghinstallation.New", zap.Error(err))
	}

	svc := service.NewService(
		github.NewClient(
			&http.Client{
				Transport: tr,
			},
		),
		"DeployPipeline",
		db.NewDynamoDB(
			dynamodb.New(sess, cfg),
			table,
		),
		pipelinestat.NewCodePipeline(
			codepipeline.New(sess, cfg),
		),
		region,
		service.SetLogger(&appLogger{}),
	)

	lambda.StartWithContext(ctx, svc.Handle)
}
