package aws

import (
	_ "errors"
	_ "strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	_ "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	_ "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

var (
	SESS *session.Session
	SVC  *dynamodb.DynamoDB
)

func InitAmazon() {
	if SESS == nil {
		SESS = session.Must(session.NewSessionWithOptions(session.Options{
			Profile: "default",
			Config: aws.Config{
				Region: aws.String("us-east-2"),
			},
			SharedConfigState: session.SharedConfigEnable,
		}))
	}
	if SVC == nil {
		SVC = dynamodb.New(SESS)
	}
}

func AddTableItem(av map[string]*dynamodb.AttributeValue, table *string) error {
	InitAmazon()
	_, err := SVC.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: table,
	})
	return err
}

func QueryTableItems(keyCond expression.KeyConditionBuilder, proj expression.ProjectionBuilder, tableName string) (*dynamodb.QueryOutput, error) {
	InitAmazon()
	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).WithProjection(proj).Build()
	if err != nil {
		return nil, err
	}
	input := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}
	return SVC.Query(input)
}

func ScanItem(keyCond *expression.ConditionBuilder, proj expression.ProjectionBuilder, tableName string) (*dynamodb.ScanOutput, error) {
	InitAmazon()
	builder := expression.NewBuilder().WithProjection(proj)
	if keyCond != nil {
		builder = builder.WithFilter(*keyCond)
	}
	expr, err := builder.Build()
	if err != nil {
		return nil, err
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}
	return SVC.Scan(params)
}
