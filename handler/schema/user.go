package schema

import "github.com/graphql-go/graphql"

type User struct {
	id       string `json:"id"`
	login    string `json:"login"`
	password string `json:"password"`
}

var graphQLUserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"login": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})
