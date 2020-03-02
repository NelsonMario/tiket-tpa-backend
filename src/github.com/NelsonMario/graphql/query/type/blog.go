package types

import "github.com/graphql-go/graphql"

var blogType *graphql.Object

func GetBlogType() *graphql.Object {
	if blogType == nil {
		blogType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "blogType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userId": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"value": &graphql.Field{
					Type: graphql.String,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
				"thumbnail": &graphql.Field{
					Type: graphql.String,
				},
				"viewer": &graphql.Field{
						Type:graphql.Int,
				},
			},
		})
	}
	return blogType
}
