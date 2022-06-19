package db

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type DynamoDB struct {
	client    *dynamodb.DynamoDB
	tableName *string

	ttl int64
}

var _ DB = (*DynamoDB)(nil)

type DynamoDBOption func(*DynamoDB)

func SetTTL(ttl int64) DynamoDBOption {
	return func(db *DynamoDB) {
		db.ttl = ttl
	}
}

func NewDynamoDB(
	client *dynamodb.DynamoDB,
	tableName string,
	opts ...DynamoDBOption,
) *DynamoDB {
	ret := &DynamoDB{
		client:    client,
		tableName: aws.String(tableName),

		ttl: 60 * 60 * 24 * 7,
	}

	for _, f := range opts {
		f(ret)
	}

	return ret
}

type dynamodbRecord struct {
	ExecutionId string `dynamodbav:"ExecutionId"`
	Timestamp   int64  `dynamodbav:"Timestamp,omitempty"`
	CheckRunId  int64  `dynamodbav:"CheckRunId,omitempty"`
	Expired     int64  `dynamodbav:"Expired,omitempty"`
}

var (
	ddbfield_executionId = "ExecutionId"
	ddbfield_timestamp   = "Timestamp"
	ddbfield_checkRunId  = "CheckRunId"
	ddbfield_expired     = "Expired"
)

func (db *DynamoDB) CreatePipeline(ctx context.Context, executionId string, checkRunId int64) error {
	now := time.Now().Unix()
	item, err := dynamodbattribute.MarshalMap(&dynamodbRecord{
		ExecutionId: executionId,
		Timestamp:   now,
		CheckRunId:  checkRunId,
		Expired:     now + db.ttl,
	})
	if err != nil {
		return err
	}

	expr, err := expression.NewBuilder().WithCondition(
		expression.Name(ddbfield_executionId).AttributeNotExists(),
	).Build()
	if err != nil {
		return err
	}

	if _, err := db.client.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: db.tableName,
		Item:      item,

		ConditionExpression:      expr.Condition(),
		ExpressionAttributeNames: expr.Names(),
	}); err != nil {
		return err
	}

	return nil
}

func (db *DynamoDB) GetCheckRunId(ctx context.Context, executionId string) (int64, error) {
	key, err := dynamodbattribute.MarshalMap(&dynamodbRecord{
		ExecutionId: executionId,
	})
	if err != nil {
		return 0, err
	}

	expr, err := expression.NewBuilder().WithProjection(
		expression.NamesList(expression.Name(ddbfield_checkRunId)),
	).Build()
	if err != nil {
		return 0, err
	}

	out, err := db.client.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: db.tableName,
		Key:       key,

		ProjectionExpression:     expr.Projection(),
		ExpressionAttributeNames: expr.Names(),
	})
	if err != nil {
		return 0, err
	}
	if len(out.Item) == 0 {
		return 0, ErrNotExists
	}

	ret := &dynamodbRecord{}
	if err := dynamodbattribute.UnmarshalMap(out.Item, ret); err != nil {
		return 0, err
	}

	return ret.CheckRunId, nil
}

func (db *DynamoDB) Scan(ctx context.Context, f func(executionId string, checkRunId int64, timestamp, expired int64)) error {
	var rerr error
	if err := db.client.ScanPagesWithContext(ctx, &dynamodb.ScanInput{
		TableName: db.tableName,
	}, func(out *dynamodb.ScanOutput, last bool) bool {
		var items []dynamodbRecord
		if err := dynamodbattribute.UnmarshalListOfMaps(out.Items, &items); err != nil {
			rerr = err
			return false
		}

		for _, item := range items {
			select {
			case <-ctx.Done():
				rerr = ctx.Err()
				return false
			default:
				f(item.ExecutionId, item.CheckRunId, item.Timestamp, item.Expired)
			}
		}

		return true
	}); err != nil {
		return err
	}

	return rerr
}

func (db *DynamoDB) ExpandExpired(ctx context.Context, executionId string, timestamp int64) error {
	now := time.Now().Unix()

	expr, err := expression.NewBuilder().WithUpdate(
		expression.Set(expression.Name(ddbfield_expired), expression.Value(now+db.ttl)).
			Set(expression.Name(ddbfield_timestamp), expression.Value(now)),
	).WithCondition(
		expression.Name(ddbfield_timestamp).AttributeNotExists().Or(
			expression.Name(ddbfield_timestamp).Equal(expression.Value(timestamp)),
		),
	).Build()
	if err != nil {
		return err
	}

	if _, err := db.client.UpdateItemWithContext(ctx, &dynamodb.UpdateItemInput{
		TableName: db.tableName,

		UpdateExpression:          expr.Update(),
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),

		Key: map[string]*dynamodb.AttributeValue{
			ddbfield_executionId: {S: aws.String(executionId)},
		},
	}); err != nil {
		return err
	}

	return nil
}
