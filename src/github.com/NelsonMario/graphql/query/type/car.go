package types

import "github.com/graphql-go/graphql"

var carType *graphql.Object

func GetCarType() *graphql.Object {
	if carType == nil {
		carType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "carType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"model": &graphql.Field{
					Type: graphql.String,
				},
				"passanger": &graphql.Field{
					Type: graphql.Int,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"vendorCar": &graphql.Field{
					Type: graphql.NewList(GetVendorCarType()),
				},
			},
		})
	}
	return carType
}
