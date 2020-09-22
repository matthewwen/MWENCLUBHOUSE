package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/stretchr/testify/assert"
	"gitlab.com/mwenclubhouse/mwenclubhouse-sdk/mwen/website/articles"
	"gitlab.com/mwenclubhouse/mwenclubhouse-sdk/mwen/website/hackathons"
	"gitlab.com/mwenclubhouse/mwenclubhouse-sdk/mwen/website/posts"
	"gitlab.com/mwenclubhouse/mwenclubhouse-sdk/mwen/website/running"
)

func TestGetAllHackathon(t *testing.T) {
	result, err := ScanItem(
		nil,
		expression.NamesList(
			expression.Name("eventName"),
			expression.Name("link"),
			expression.Name("name"),
		),
		"MWENWEBSITE-Hackathons",
	)
	assert.Equal(t, err, nil)
	hackathons := make([]hackathons.Event, len(result.Items))
	for i, item := range result.Items {
		if err == nil {
			err = dynamodbattribute.UnmarshalMap(item, &hackathons[i])
		}
	}
	fmt.Println(hackathons)
}

func TestGetAllArticle(t *testing.T) {
	var tableName string = "MWENWEBSITE-Page-View-Articles"
	result, err := ScanItem(
		nil,
		expression.NamesList(
			expression.Name("ref"),
			expression.Name("description"),
			expression.Name("service"),
			expression.Name("title"),
		),
		tableName,
	)
	assert.Equal(t, err, nil)
	a := make([]articles.Articles, len(result.Items))
	for i, item := range result.Items {
		if err == nil {
			err = dynamodbattribute.UnmarshalMap(item, &a[i])
		}
	}
}

func TestGetRunningData(t *testing.T) {
	KeyExpression := expression.Name("date").GreaterThanEqual(expression.Value(-1))
	result, err := ScanItem(
		&KeyExpression,
		expression.NamesList(
			expression.Name("date"),
			expression.Name("RunningDistance"),
			expression.Name("RunningTime"),
			expression.Name("StepCount"),
		),
		"MWENWEBSITE-Running",
	)
	assert.Equal(t, err, nil)
	response := make([]running.Data, *result.Count)
	for i, item := range result.Items {
		if err == nil {
			err = dynamodbattribute.UnmarshalMap(item, &response[i])
		}
	}
	fmt.Println(response)
}

func TestGetPostDataIsSuccessful(t *testing.T) {
	result, err := ScanItem(
		nil,
		expression.NamesList(
			expression.Name("ref"),
			expression.Name("flares"),
			expression.Name("p"),
			expression.Name("title"),
		),
		"MWENWEBSITE-POSTS",
	)
	assert.Equal(t, err, nil)
	unParsedPosts := make([]posts.UnParsedPost, len(result.Items))
	for i, item := range result.Items {
		if err == nil {
			err = dynamodbattribute.UnmarshalMap(item, &unParsedPosts[i])
		}
	}
	fmt.Println(unParsedPosts)
}
