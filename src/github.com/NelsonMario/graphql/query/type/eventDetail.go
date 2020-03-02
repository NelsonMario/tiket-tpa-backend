package types

import "github.com/graphql-go/graphql"

var eventDetailType *graphql.Object

func GetEventDetailType() *graphql.Object {
	if eventDetailType == nil {
		eventDetailType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "eventDetailType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return eventDetailType
}
