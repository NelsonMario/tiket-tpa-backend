package types

import "github.com/graphql-go/graphql"

var sliderType *graphql.Object

func GetSliderType() *graphql.Object {
	if sliderType == nil {
		sliderType = graphql.NewObject(graphql.ObjectConfig{
			Name: "SliderType",
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
	return sliderType
}
