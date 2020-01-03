package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllSlider(p graphql.ResolveParams) (i interface{}, e error) {
	slider, err := models.GetAllSlider()
	return slider, err
}

func InsertSlider(p graphql.ResolveParams) (i interface{}, e error) {
	name := p.Args["name"].(string)

	slider, err := models.InsertSlider(name)

	if err != nil {
		return nil, err
	}
	return slider, nil
}