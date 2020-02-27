package types

import "github.com/graphql-go/graphql"

var trainType *graphql.Object

func GetTrainType() *graphql.Object{
	if trainType == nil {
		trainType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "trainType",
			Fields:      graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
			IsTypeOf:    nil,
			Description: "",
		})

	}
	return trainType
}

