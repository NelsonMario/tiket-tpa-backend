package types

import "github.com/graphql-go/graphql"

var reviewType *graphql.Object

func GetReviewType() *graphql.Object {
	if reviewType == nil {
		reviewType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "reviewType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"review": &graphql.Field{
					Type: graphql.Int,
				},
				"user": &graphql.Field{
					Type: graphql.String,
				},
				"category": &graphql.Field{
					Type: graphql.String,
				},
				"caption": &graphql.Field{
					Type: graphql.String,
				},
				"hotelRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"createdAt": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		})
	}
	return reviewType
}

