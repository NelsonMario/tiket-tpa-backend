package types

import "github.com/graphql-go/graphql"

var roomType *graphql.Object

func GetRoomType() *graphql.Object {
	if roomType == nil {
		roomType = graphql.NewObject(graphql.ObjectConfig{
			Name: "RoomType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"checkIn": &graphql.Field{
					Type: graphql.DateTime,
				},
				"checkOut": &graphql.Field{
					Type: graphql.DateTime,
				},
				"bed": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"maxGuest": &graphql.Field{
					Type: graphql.Int,
				},
				"size": &graphql.Field{
					Type: graphql.Int,
				},
				"roomFacility": &graphql.Field{
					Type: graphql.NewList(GetRoomFacilityType()),
				},
				"type": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return roomType
}
