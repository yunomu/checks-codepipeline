package stat

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/google/subcommands"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codepipeline"

	"github.com/yunomu/checks-codepipeline/lambda/checks/pipelinestat"
)

type Command struct {
	region *string

	name   *string
	execId *string
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "stat" }
func (c *Command) Synopsis() string { return "pipeline stat" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.SetOutput(os.Stderr)

	c.region = f.String("region", "ap-northeast-1", "Endpoint of DynamoDB (default: config.json)")
	c.name = f.String("name", "", "Pipeline Name")
	c.execId = f.String("execId", "", "Pipeline Execution ID")
}

func (c *Command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	sess, err := session.NewSession(aws.NewConfig().WithRegion(*c.region))
	if err != nil {
		log.Fatalf("session.NewSession: %v", err)
	}

	client := pipelinestat.NewCodePipeline(codepipeline.New(sess))

	stat, err := client.GetStat(ctx, *c.name, *c.execId)
	if err != nil {
		log.Fatalf("GetStat: %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(stat); err != nil {
		log.Fatalf("Encode: %v", err)
	}

	return subcommands.ExitSuccess
}
