package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllStation(p graphql.ResolveParams) (i interface{}, e error) {
	station, err := models.GetAllStation()
	return station, err
}

func GetDistinctStation(p graphql.ResolveParams) (i interface{}, e error) {
	station, err := models.GetDistinctStation()
	return station, err
}

func InsertStation(p graphql.ResolveParams)(i interface{}, e error){

	code := p.Args["code"].(string)
	name := p.Args["name"].(string)
	city := p.Args["city"].(string)
	cityCode := p.Args["cityCode"].(string)
	province := p.Args["province"].(string)
	country := p.Args["country"].(string)


	station, err := models.InsertStation(code, name, city, cityCode, province, country)

	if err != nil {
		return nil, err
	}
	return station, nil
}

