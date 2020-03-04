package types

import "github.com/graphql-go/graphql"

var tripReviewType *graphql.Object

func GetTripReviewType() *graphql.Object {
	if tripReviewType == nil {
		tripReviewType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "tripReviewType",
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
	return tripReviewType
}

