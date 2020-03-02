package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func UpdateEventDetail(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	name := p.Args["name"].(string)
	eventRefer := p.Args["eventRefer"].(int)
	price := p.Args["price"].(int)

	eventDetail, err := models.UpdateEventDetail(id, name, eventRefer, price)

	if err != nil {
		return nil, err
	}
	return eventDetail, nil
}


func RemoveEventDetail(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)

	eventDetail, err := models.RemoveEventDetail(id)

	if err != nil {
		return nil, err
	}
	return eventDetail, nil
}

func InsertEventDetail(p graphql.ResolveParams)(i interface{}, e error){

	name := p.Args["name"].(string)
	eventRefer := p.Args["eventRefer"].(int)
	price := p.Args["price"].(int)

	eventDetail, err := models.InsertEventDetail(name, eventRefer, price)

	if err != nil {
		return nil, err
	}
	return eventDetail, nil
}

