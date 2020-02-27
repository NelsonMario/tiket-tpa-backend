package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllRoom(p graphql.ResolveParams)(i interface{}, e error){
	room, err :=  models.GetAllRoom()
	return room, err
}


func GetRoomBySchedule(p graphql.ResolveParams) (i interface{}, e error){
	location := p.Args["location"].(int)
	from := p.Args["fromDate"].(string)
	to := p.Args["toDate"].(string)

	room, err := models.GetRoomBySchedule(location, from, to)

	if err != nil {
		return nil, err
	}
	return room, nil
}

