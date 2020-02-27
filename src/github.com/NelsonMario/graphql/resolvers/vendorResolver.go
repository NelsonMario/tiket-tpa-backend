package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllVendor(p graphql.ResolveParams)(i interface{}, e error){
	vendor, err :=  models.GetAllVendor()
	return vendor, err
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
