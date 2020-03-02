package types

import "github.com/graphql-go/graphql"

var promoDetailType *graphql.Object

func getPromoDetailType() *graphql.Object {
	if promoDetailType == nil {
		promoDetailType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "promoDetailType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"disc": &graphql.Field{
					Type: graphql.Int,
				},
				"promoCode": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return promoDetailType
}
