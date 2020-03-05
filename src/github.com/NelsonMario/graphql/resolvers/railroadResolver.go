package resolvers
import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
	"time"
)

func GetAllRailroad(p graphql.ResolveParams) (i interface{}, e error) {
	railroad, err := models.GetAllRailroad()
	return railroad, err
}

func InsertRailroad(p graphql.ResolveParams)(i interface{}, e error){
	//notes : add flight routes
	trainRefer := p.Args["trainRefer"].(int)
	fromRefer := p.Args["fromRefer"].(int)
	toRefer := p.Args["toRefer"].(int)
	departure := p.Args["departure"].(time.Time)
	arrival := p.Args["arrival"].(time.Time)
	duration := p.Args["duration"].(int)
	price := p.Args["price"].(int)
	tax := p.Args["tax"].(int)
	serviceCharge := p.Args["serviceCharge"].(int)

	if(trainRefer == -1 || fromRefer == -1 || toRefer == -1 || duration == -1 || price == -1 || tax == -1 || serviceCharge == -1){
		return 1, nil
	}

	railroad, err := models.InsertRailroad(trainRefer, fromRefer, toRefer, departure, arrival, duration, price, tax ,serviceCharge)

	if err != nil {
		return nil, err
	}
	return railroad, nil
}

func GetRailroadByToAndFrom(p graphql.ResolveParams)(i interface{}, e error){
	from := p.Args["from"].(string)
	to := p.Args["to"].(string)

	railroad, err := models.GetRailroadByFromAndTo(from, to)

	if err != nil {
		return nil, err
	}
	return railroad, nil
}

func GetRailroadByOneSchedule(p graphql.ResolveParams)(i interface{}, e error){
	from := p.Args["from"].(string)
	to := p.Args["to"].(string)
	date := p.Args["date"].(string)

	railroad, err := models.GetRailroadByOneSchedule(from, to, date)

	if err != nil {
		return nil, err
	}
	return railroad, nil
}

func GetRailroadBySchedule(p graphql.ResolveParams)(i interface{}, e error){
	from := p.Args["from"].(string)
	to := p.Args["to"].(string)
	arrival := p.Args["arrival"].(string)
	departure := p.Args["departure"].(string)

	railroad, err := models.GetRailroadBySchedule(from, to, departure, arrival)

	if err != nil {
		return nil, err
	}
	return railroad, nil
}
