package main

import (
	"context"
	"encoding/json"

	"serverless-go-graphql/handler/schema"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
)

type RequestBody struct {
	Query     string `json:"query"`
	Variables struct {
	} `json:"variables"`
	OperationName string `json:"operationName"`
}

func executeQuery(request RequestBody, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: request.Query,
		OperationName: request.OperationName,
	})

	return result
}

// Handler of HTTP event
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	requestBody := RequestBody{}
	err := json.Unmarshal([]byte(request.Body), &requestBody)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
	}

	graphQLResult := executeQuery(requestBody, schema.Schema)
	responseJSON, err := json.Marshal(graphQLResult)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, err
	}

	return events.APIGatewayProxyResponse{Body: string(responseJSON[:]), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
