package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
	"time"
)

func GetAllFlight(p graphql.ResolveParams) (i interface{}, e error) {
	airports, err := models.GetAllFlight()
	return airports, err
}

func InsertFlight(p graphql.ResolveParams)(i interface{}, e error){
	//notes : add flight routes


	airlineRefer := p.Args["airlineRefer"].(int)
	fromRefer := p.Args["fromRefer"].(int)
	toRefer := p.Args["toRefer"].(int)
	departure := p.Args["departure"].(time.Time)
	arrival := p.Args["arrival"].(time.Time)
	duration := p.Args["duration"].(int)
	price := p.Args["price"].(int)
	tax := p.Args["tax"].(int)
	serviceCharge := p.Args["serviceCharge"].(int)

	if(price == -1 || duration == -1 || tax == -1 || serviceCharge == -1){
		return 1, nil
	}


	flight, err := models.InsertFlight(airlineRefer, fromRefer, toRefer, departure, arrival, duration, price, tax ,serviceCharge)

	if err != nil {
		return nil, err
	}
	return flight, nil
}

func GetFlightByToAndFrom(p graphql.ResolveParams)(i interface{}, e error){
	from := p.Args["from"].(string)
	to := p.Args["to"].(string)

	flight, err := models.GetFlightByFromAndTo(from, to)

	if err != nil {
		return nil, err
	}
	return flight, nil
}


func GetFlightByOneSchedule(p graphql.ResolveParams)(i interface{}, e error){
	from := p.Args["from"].(string)
	to := p.Args["to"].(string)
	date := p.Args["date"].(string)
	flight, err := models.GetFlightByOneSchedule(from, to, date)

	if err != nil {
		return nil, err
	}
	return flight, nil
}


func GetFlightBySchedule(p graphql.ResolveParams)(i interface{}, e error){
	from := p.Args["from"].(string)
	to := p.Args["to"].(string)
	arrival := p.Args["arrival"].(string)
	departure := p.Args["departure"].(string)

	flight, err := models.GetFlightBySchedule(from, to, departure, arrival)

	if err != nil {
		return nil, err
	}
	return flight, nil
}

func UpdateFlight(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	airline := p.Args["airline_refer"].(int)
	from := p.Args["from_refer"].(int)
	to := p.Args["to_refer"].(int)
	arrival := p.Args["arrival"].(time.Time)
	departure := p.Args["departure"].(time.Time)
	duration := p.Args["duration"].(int)
	price := p.Args["price"].(int)
	tax := p.Args["tax"].(int)
	serviceCharge := p.Args["service_charge"].(int)


	flight, err := models.UpdateFlight(id, airline, from, to, arrival, departure, duration, price, tax, serviceCharge)

	if err != nil {
		return nil, err
	}
	return flight, nil
}


func RemoveFlight(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)

	flight, err := models.RemoveFlight(id)

	if err != nil {
		return nil, err
	}
	return flight, nil
}