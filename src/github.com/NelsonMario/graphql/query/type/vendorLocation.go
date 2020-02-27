package types

import "github.com/graphql-go/graphql"

var vendorLocationType *graphql.Object

func GetVendorLocationType() *graphql.Object {
	if vendorLocationType == nil {
		vendorLocationType = graphql.NewObject(graphql.ObjectConfig{
			Name: "vendorLocationType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"location": &graphql.Field{
					Type: GetLocationType(),
				},
				"LocationRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"VendorLocationRefer": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return vendorLocationType
}
