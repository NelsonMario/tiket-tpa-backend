package types

import "github.com/graphql-go/graphql"

var flightType *graphql.Object

func GetFlightType() *graphql.Object {
	if flightType == nil {
		flightType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "FlightType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"airline": &graphql.Field{
					Type: GetAirlineType(),
				},
				"airlineRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"from": &graphql.Field{
					Type: GetAirportType(),
				},
				"fromRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"to": &graphql.Field{
					Type: GetAirportType(),
				},
				"toRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"flightRoutes": &graphql.Field{
					Type: graphql.NewList(GetFlightRoutesType()),
				},
				"flightFacility": &graphql.Field{
					Type: graphql.NewList(GetFlightFacilityType()),
				},
				"departure": &graphql.Field{
					Type: graphql.DateTime,
				},
				"arrival": &graphql.Field{
					Type: graphql.DateTime,
				},
				"duration": &graphql.Field{
					Type: graphql.Int,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"tax": &graphql.Field{
					Type: graphql.Int,
				},
				"serviceCharge": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return flightType
}
