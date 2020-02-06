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
	airlineRefer := p.Args["airlineRefer"].(uint)
	fromRefer := p.Args["fromRefer"].(uint)
	toRefer := p.Args["toRefer"].(uint)
	departure := p.Args["departure"].(time.Time)
	arrival := p.Args["arrival"].(time.Time)
	duration := p.Args["duration"].(uint)
	price := p.Args["price"].(int)
	tax := p.Args["tax"].(int)
	serviceCharge := p.Args["serviceCharge"].(int)


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
