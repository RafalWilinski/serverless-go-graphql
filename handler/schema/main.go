package schema

import (
	"serverless-go-graphql/handler/services"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getUser": &graphql.Field{
			Type:        graphQLUserType,
			Description: "Get single user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				user := User{}
				id, _ := params.Args["id"].(string)
				result, err := services.GetUser(id)

				if err != nil {
					return nil, err
				}

				err = dynamodbattribute.UnmarshalMap(result.Item, &user)
				return user, err
			},
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: graphql.String,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return "hello", nil
			},
		},
	},
})

// Schema - GraphQL root schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
