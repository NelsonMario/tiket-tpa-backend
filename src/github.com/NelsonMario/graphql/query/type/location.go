package types

import "github.com/graphql-go/graphql"

var locationType *graphql.Object

func GetLocationType() *graphql.Object {
	if locationType == nil {
		locationType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "LocationType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"city": &graphql.Field{
					Type: graphql.String,
				},
				"country": &graphql.Field{
					Type: graphql.String,
				},
				"lat": &graphql.Field{
					Type: graphql.Float,
				},
				"lng": &graphql.Field{
					Type: graphql.Float,
				},
			},
		})
	}
	return locationType
}
