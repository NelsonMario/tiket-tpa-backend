package types

import "github.com/graphql-go/graphql"

var flightFacilityType *graphql.Object

func GetFlightFacilityType() *graphql.Object {
	if flightFacilityType == nil {
		flightFacilityType = graphql.NewObject(graphql.ObjectConfig{
			Name: "flightFacilityType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"facility": &graphql.Field{
					Type: GetFacilityType(),
				},
				"facilityRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"flightReferFacility": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return flightFacilityType
}
