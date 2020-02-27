package types

import "github.com/graphql-go/graphql"

var hotelFacilityType *graphql.Object

func GetHotelFacilityType() *graphql.Object {
	if hotelFacilityType == nil {
		hotelFacilityType = graphql.NewObject(graphql.ObjectConfig{
			Name: "hotelFacilityType",
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
				"hotelReferFacility": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return hotelFacilityType
}
