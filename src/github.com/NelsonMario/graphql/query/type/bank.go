package types

import "github.com/graphql-go/graphql"

var bankType *graphql.Object

func GetBankType() *graphql.Object{
	if bankType == nil {
		bankType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "bankType",
			Fields:      graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"desc": &graphql.Field{
					Type: graphql.String,
				},
			},
			IsTypeOf:    nil,
			Description: "",
		})

	}
	return bankType
}