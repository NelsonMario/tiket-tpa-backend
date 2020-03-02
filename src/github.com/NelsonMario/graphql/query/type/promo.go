package types

import "github.com/graphql-go/graphql"

var promoType *graphql.Object

func GetPromoType() *graphql.Object {
	if promoType == nil {
		promoType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "promoType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"startDate": &graphql.Field{
					Type: graphql.DateTime,
				},
				"endDate": &graphql.Field{
					Type: graphql.DateTime,
				},
				"promoDetail": &graphql.Field{
					Type: graphql.NewList(getPromoDetailType()),
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return promoType
}

