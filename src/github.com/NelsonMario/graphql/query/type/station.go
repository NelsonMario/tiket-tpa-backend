package types

import "github.com/graphql-go/graphql"

var stationType *graphql.Object

func GetStationType() *graphql.Object {
	if stationType == nil {
		stationType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "stationType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"code": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"city": &graphql.Field{
					Type: graphql.String,
				},
				"cityCode": &graphql.Field{
					Type: graphql.String,
				},
				"province": &graphql.Field{
					Type: graphql.String,
				},
				"country": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return stationType
}
