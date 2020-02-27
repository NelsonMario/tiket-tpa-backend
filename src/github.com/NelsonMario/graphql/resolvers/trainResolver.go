package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
	"time"
)

func GetAllTrain(p graphql.ResolveParams)(i interface{}, e error){
	train, err :=  models.GetAllTrain()
	return train, err
}

func GetTrainById(p graphql.ResolveParams)(i interface{}, e error){
	id := p.Args["id"].(int)
	train, err := models.GetTrainById(id)
	return train, err
}

func InsertTrain(p graphql.ResolveParams)(i interface{}, e error){
	name := p.Args["name"].(string)

	train, err := models.InsertTrain(name)

	if err != nil {
		return nil, err
	}

	return train, e
}

func UpdateRailroad(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	train := p.Args["train_refer"].(int)
	from := p.Args["from_refer"].(int)
	to := p.Args["to_refer"].(int)
	arrival := p.Args["arrival"].(time.Time)
	departure := p.Args["departure"].(time.Time)
	duration := p.Args["duration"].(int)
	price := p.Args["price"].(int)
	tax := p.Args["tax"].(int)
	serviceCharge := p.Args["service_charge"].(int)


	railroad, err := models.UpdateRailroad(id, train, from, to, arrival, departure, duration, price, tax, serviceCharge)

	if err != nil {
		return nil, err
	}
	return railroad, nil
}


func RemoveRailroad(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)

	railroad, err := models.RemoveRailroad(id)

	if err != nil {
		return nil, err
	}
	return railroad, nil
}
