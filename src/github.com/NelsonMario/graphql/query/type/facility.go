package types

import "github.com/graphql-go/graphql"

var facilityType *graphql.Object

func GetFacilityType() *graphql.Object {
	if facilityType == nil {
		facilityType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "facilityType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return facilityType
}
