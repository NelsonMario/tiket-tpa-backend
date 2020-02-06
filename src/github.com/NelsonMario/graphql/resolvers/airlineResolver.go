package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllAirline(p graphql.ResolveParams)(i interface{}, e error){
	airline, err :=  models.GetAllAirline()
	return airline, err
}

func GetAirlineById(p graphql.ResolveParams)(i interface{}, e error){
	id := p.Args["id"].(int)
	airline, err := models.GetAirlineById(id)
	return airline, err
}

func InsertAirline(p graphql.ResolveParams)(i interface{}, e error){
	name := p.Args["name"].(string)

	airline, err := models.InsertAirline(name)

	if err != nil {
		return nil, err
	}

	return airline, e
}
