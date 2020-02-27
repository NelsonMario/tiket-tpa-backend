package types

import "github.com/graphql-go/graphql"

var vendorCarType *graphql.Object

func GetVendorCarType() *graphql.Object {
	if vendorCarType == nil {
		vendorCarType = graphql.NewObject(graphql.ObjectConfig{
			Name: "vendorCarType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"vendor": &graphql.Field{
					Type: GetVendorType(),
				},
				"vendorRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"VendorCarRefer": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return vendorCarType
}
