package types

import "github.com/graphql-go/graphql"

var eventType *graphql.Object

func GetEventType() *graphql.Object {
	if eventType == nil {
		eventType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "eventType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"startDate": &graphql.Field{
					Type: graphql.DateTime,
				},
				"endDate": &graphql.Field{
					Type: graphql.DateTime,
				},
				"location": &graphql.Field{
					Type: GetLocationType(),
				},
				"eventLat": &graphql.Field{
				Type: graphql.Float,
				},
				"eventLng": &graphql.Field{
					Type: graphql.Float,
				},
				"eventDetail": &graphql.Field{
					Type: graphql.NewList(GetEventDetailType()),
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"category": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return eventType
}

