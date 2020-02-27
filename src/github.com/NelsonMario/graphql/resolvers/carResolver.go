package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllCar(p graphql.ResolveParams)(i interface{}, e error){
	car, err :=  models.GetAllCar()
	return car, err
}

func GetCarByLocation(p graphql.ResolveParams)(i interface{}, e error){
	location := p.Args["location"].(string)
	car, err := models.GetCarByLocation(location)
	return car, err
}

//func GetDistinctHotel(p graphql.ResolveParams) (i interface{}, e error) {
//	hotel, err := models.GetDistinctHotel()
//	return hotel, err
//}
//
//func InsertHotel(p graphql.ResolveParams)(i interface{}, e error){
//	name := p.Args["name"].(string)
//	rating := p.Args["rating"].(int)
//	_type := p.Args["type"].(string)
//	city := p.Args["city"].(string)
//	cityCode := p.Args["cityCode"].(string)
//	province := p.Args["province"].(string)
//	country := p.Args["country"].(string)
//
//	hotel, err := models.InsertHotel(name, rating, _type, city, cityCode, province, country)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return hotel, e
//}
