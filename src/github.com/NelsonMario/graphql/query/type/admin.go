package types

import "github.com/graphql-go/graphql"

var adminType *graphql.Object

func GetAdminType() *graphql.Object {
	if adminType == nil {
		adminType = graphql.NewObject(graphql.ObjectConfig{
			Name: "AdminType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return adminType
}
