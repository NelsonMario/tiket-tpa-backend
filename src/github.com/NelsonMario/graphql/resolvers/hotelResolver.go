package resolvers

import (
	"fmt"
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllHotel(p graphql.ResolveParams)(i interface{}, e error){
	hotel, err :=  models.GetAllHotel()
	return hotel, err
}

func GetHotelById(p graphql.ResolveParams)(i interface{}, e error){
	id := p.Args["id"].(int)

	hotel, err :=  models.GetHotelById(id)
	return hotel, err
}

func GetHotelByLocation(p graphql.ResolveParams)(i interface{}, e error){
	location := p.Args["location"].(string)

	hotel, err :=  models.GetNearestHotelByLocation(location)
	return hotel, err
}

func GetNearestHotelByLocation(p graphql.ResolveParams)(i interface{}, e error){
	location := p.Args["location"].(string)

	hotel, err :=  models.GetNearestHotelByLocation(location)
	return hotel, err
}

func GetDistinctHotel(p graphql.ResolveParams) (i interface{}, e error) {
	hotel, err := models.GetDistinctHotel()
	return hotel, err
}

func InsertHotel(p graphql.ResolveParams)(i interface{}, e error){
	name := p.Args["name"].(string)
	rating := p.Args["rating"].(int)
	_type := p.Args["type"].(string)
	locationRefer := p.Args["locationRefer"].(int)
	hotelLat := p.Args["hotelLat"].(float64)
	hotelLng := p.Args["hotelLng"].(float64)

	hotel, err := models.InsertHotel(name, rating, _type, locationRefer, hotelLat, hotelLng)

	if err != nil {
		return nil, err
	}

	return hotel, e
}

func UpdateHotel(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	name := p.Args["name"].(string)
	rating := p.Args["rating"].(int)


	hotel, err := models.UpdateHotel(id, name, rating)

	if err != nil {
		return nil, err
	}
	return hotel, nil
}

func RemoveHotel(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)

	hotel, err := models.RemoveHotel(id)
	fmt.Println(id)
	if err != nil {
		return nil, err
	}
	return hotel, nil
}

