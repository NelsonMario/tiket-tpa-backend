package types

import "github.com/graphql-go/graphql"

var userType *graphql.Object

func GetUserType() *graphql.Object {
	if userType == nil {
		userType = graphql.NewObject(graphql.ObjectConfig{
			Name: "UserType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"firstName": &graphql.Field{
					Type: graphql.String,
				},
				"lastName": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"phoneNumber": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
				"cityName": &graphql.Field{
					Type: graphql.String,
				},
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"postcode": &graphql.Field{
					Type: graphql.String,
				},
				"language": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return userType
}
