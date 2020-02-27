package types

import "github.com/graphql-go/graphql"

var hotelType *graphql.Object

func GetHotelType() *graphql.Object {
	if hotelType == nil {
		hotelType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "HotelType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"rating": &graphql.Field{
					Type: graphql.Int,
				},
				"type": &graphql.Field{
					Type: graphql.String,
				},
				"hotelFacility": &graphql.Field{
					Type: graphql.NewList(GetHotelFacilityType()),
				},
				"location": &graphql.Field{
					Type: GetLocationType(),
				},
				"hotelRoom": &graphql.Field{
					Type: graphql.NewList(GetHotelRoomType()),
				},
				"hotelLat": &graphql.Field{
					Type: graphql.Float,
				},
				"hotelLng": &graphql.Field{
					Type: graphql.Float,
				},
			},
		})
	}
	return hotelType
}
