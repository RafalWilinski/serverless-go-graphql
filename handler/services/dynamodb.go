package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createSession() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	return dynamodb.New(sess)
}

func GetUser(id string) (*dynamodb.GetItemOutput, error) {
	svc := createSession()

	return svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("go-graphql-users"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
}
