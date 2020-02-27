package types

import "github.com/graphql-go/graphql"

var railroadRoutesType *graphql.Object

func GetRailroadRoutesType() *graphql.Object {
	if railroadRoutesType == nil {
		railroadRoutesType = graphql.NewObject(graphql.ObjectConfig{
			Name: "railroadRoutesType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"station": &graphql.Field{
					Type: GetStationType(),
				},
				"stationRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"railroadReferRoute": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return railroadRoutesType
}
