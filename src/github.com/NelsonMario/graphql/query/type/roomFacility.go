package types

import "github.com/graphql-go/graphql"

var roomFacilityType *graphql.Object

func GetRoomFacilityType() *graphql.Object {
	if roomFacilityType == nil {
		roomFacilityType = graphql.NewObject(graphql.ObjectConfig{
			Name: "roomFacilityType",
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
				"facilityType": &graphql.Field{
					Type: graphql.String,
				},
				"roomReferFacility": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return roomFacilityType
}
