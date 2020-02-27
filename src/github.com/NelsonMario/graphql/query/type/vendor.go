package types

import "github.com/graphql-go/graphql"

var vendorType *graphql.Object

func GetVendorType() *graphql.Object {
	if vendorType == nil {
		vendorType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "VendorType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"vendorLocation": &graphql.Field{
					Type: graphql.NewList(GetVendorLocationType()),
				},
			},
		})
	}
	return vendorType
}
