package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllAirport(p graphql.ResolveParams) (i interface{}, e error) {
	airports, err := models.GetAllAirport()
	return airports, err
}

func GetDistinctAirport(p graphql.ResolveParams) (i interface{}, e error) {
	airports, err := models.GetDistinctAirport()
	return airports, err
}

func InsertAirport(p graphql.ResolveParams)(i interface{}, e error){

	code := p.Args["code"].(string)
	name := p.Args["name"].(string)
	city := p.Args["city"].(string)
	cityCode := p.Args["cityCode"].(string)
	province := p.Args["province"].(string)
	country := p.Args["country"].(string)


	airport, err := models.InsertAirport(code, name, city, cityCode, province, country)

	if err != nil {
		return nil, err
	}
	return airport, nil
}

