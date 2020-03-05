package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
	"time"
)

func GetAllEvent(p graphql.ResolveParams)(i interface{}, e error){
	event, err :=  models.GetAllEvent()
	return event, err
}



func GetNearestEventByLocation(p graphql.ResolveParams)(i interface{}, e error){
	location := p.Args["location"].(string)

	event, err :=  models.GetNearestEventByLocation(location)
	return event, err
}

func GetEventByCategory(p graphql.ResolveParams)(i interface{}, e error){
	category := p.Args["category"].(string)

	event, err :=  models.GetAllEventByCategory(category)
	return event, err
}

func GetEventByLocation(p graphql.ResolveParams)(i interface{}, e error){
	location := p.Args["location"].(string)

	event, err :=  models.GetEventByLocation(location)
	return event, err
}

func UpdateEvent(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	name := p.Args["name"].(string)
	locationRefer := p.Args["locationRefer"].(int)
	startDate := p.Args["startDate"].(time.Time)
	endDate := p.Args["endDate"].(time.Time)

	event, err := models.UpdateEvent(id, name, locationRefer, startDate, endDate)

	if err != nil {
		return nil, err
	}
	return event, nil
}

func GetEventById(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)

	event, err := models.GetEventById(id)

	if err != nil {
		return nil, err
	}
	return event, nil
}

func RemoveEvent(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)

	event, err := models.RemoveEvent(id)

	if err != nil {
		return nil, err
	}
	return event, nil
}

func InsertEvent(p graphql.ResolveParams)(i interface{}, e error){

	name := p.Args["name"].(string)
	locationRefer := p.Args["locationRefer"].(int)
	startDate := p.Args["startDate"].(time.Time)
	endDate := p.Args["endDate"].(time.Time)
	eventLat := p.Args["eventLat"].(float64)
	eventLng := p.Args["eventLng"].(float64)
	category := p.Args["category"].(string)
	description := p.Args["description"].(string)

	if(name == " " || locationRefer == -1 || category == " " || description == " "){
		return 1, nil
	}

	event, err := models.InsertEvent(name, locationRefer, startDate, endDate, eventLat, eventLng, category, description)

	if err != nil {
		return nil, err
	}
	return event, nil
}

