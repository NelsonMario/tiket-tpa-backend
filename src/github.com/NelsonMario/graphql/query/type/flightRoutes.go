package types

import "github.com/graphql-go/graphql"

var flightRoutesType *graphql.Object

func GetFlightRoutesType() *graphql.Object {
	if flightRoutesType == nil {
		flightRoutesType = graphql.NewObject(graphql.ObjectConfig{
			Name: "flightRoutesType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"airport": &graphql.Field{
					Type: GetAirportType(),
				},
				"airlineRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"flightReferRoute": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return flightRoutesType
}
