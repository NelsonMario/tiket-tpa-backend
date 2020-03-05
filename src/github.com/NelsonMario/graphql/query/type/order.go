package types

import "github.com/graphql-go/graphql"

var orderType *graphql.Object

func GetOrderType() *graphql.Object {
	if orderType == nil {
		orderType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "orderType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"hotelRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"checkIn": &graphql.Field{
					Type: graphql.DateTime,
				},
				"checkOut": &graphql.Field{
					Type: graphql.DateTime,
				},
				"totalGuest": &graphql.Field{
					Type: graphql.Int,
				},
				"confirmationCode": &graphql.Field{
					Type: graphql.String,
				},
				"totalNightFee": &graphql.Field{
					Type: graphql.Int,
				},
				"cleaningFee": &graphql.Field{
					Type: graphql.Int,
				},
				"serviceFee": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return orderType
}

