package types

import "github.com/graphql-go/graphql"

var airlineType *graphql.Object

func GetAirlineType() *graphql.Object{
	if airlineType == nil {
		airlineType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "AirlineType",
			Fields:      graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
			IsTypeOf:    nil,
			Description: "",
		})

	}
	return airlineType
}

