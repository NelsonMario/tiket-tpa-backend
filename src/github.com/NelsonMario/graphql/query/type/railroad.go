package types

import "github.com/graphql-go/graphql"

var railroadType *graphql.Object

func GetRailroadType() *graphql.Object {
	if railroadType == nil {
		railroadType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "railroadType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"trainRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"train": &graphql.Field{
					Type: GetTrainType(),
				},
				"stationRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"from": &graphql.Field{
					Type: GetStationType(),
				},
				"fromRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"to": &graphql.Field{
					Type: GetStationType(),
				},
				"toRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"railroadRoutes": &graphql.Field{
					Type: graphql.NewList(GetRailroadRoutesType()),
				},
				"departure": &graphql.Field{
					Type: graphql.DateTime,
				},
				"arrival": &graphql.Field{
					Type: graphql.DateTime,
				},
				"duration": &graphql.Field{
					Type: graphql.Int,
				},
				"class": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"tax": &graphql.Field{
					Type: graphql.Int,
				},
				"serviceCharge": &graphql.Field{
					Type: graphql.Int,
				},

			},
		})
	}
	return railroadType
}
