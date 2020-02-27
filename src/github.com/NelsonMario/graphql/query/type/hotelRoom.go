package types

import "github.com/graphql-go/graphql"

var hotelRoomType *graphql.Object

func GetHotelRoomType() *graphql.Object {
	if hotelRoomType == nil {
		hotelRoomType = graphql.NewObject(graphql.ObjectConfig{
			Name: "hotelRoomType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"room": &graphql.Field{
					Type: GetRoomType(),
				},
				"roomRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"hotelReferRoom": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return hotelRoomType
}
