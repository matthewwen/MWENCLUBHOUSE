package todoplanner

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	mwenAWS "gitlab.com/mwenclubhouse/mwenclubhouse-sdk/mwen/aws"
)

type Task struct {
	Ref       string `json:"ref"`
	Name      string `json:"name"`
	Priority  int    `json:"priority"`
	Completed bool   `json:"completed"`
	Due       int    `json:"due"`
	Type      int    `json:"type"`
	Order     int    `json:"order"`
	Parent    string `json:"parent"`
}

var (
	TableName = "TodoPlanner"
)

func AddTask(s string) error {
	var av map[string]*dynamodb.AttributeValue
	var err error

	item := Task{Ref: "", Name: "", Priority: 5, Completed: false, Due: -1, Type: 1, Order: 1, Parent: "/"}

	err = json.Unmarshal([]byte(s), &item)
	if err != nil {
		return err
	}

	if item.Name == "" {
		err = errors.New("name is not set")
		return err
	}

	av, err = dynamodbattribute.MarshalMap(item)
	if err == nil && item.Ref == "" {
		item.Ref = "putref"
	}

	if err == nil {
		err = mwenAWS.AddTableItem(av, &TableName)
	}
	return err
}

func GetTask(ref string) (Task, error) {
	task := Task{}

	result, err := mwenAWS.QueryTableItems(
		expression.Key("ref").Equal(expression.Value(ref)),
		expression.NamesList(
			expression.Name("ref"),
			expression.Name("name"),
			expression.Name("priority"),
			expression.Name("completed"),
			expression.Name("due"),
			expression.Name("type"),
			expression.Name("order"),
			expression.Name("parent"),
		),
		"TodoPlanner",
	)

	if err == nil && (result == nil || *result.Count != 1) {
		err = errors.New("result is not found")
	}

	if err == nil {
		err = dynamodbattribute.UnmarshalMap(result.Items[0], &task)
	}

	return task, err
}
