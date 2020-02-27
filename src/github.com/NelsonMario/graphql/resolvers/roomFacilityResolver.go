package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllRoomFacility(p graphql.ResolveParams) (i interface{}, e error) {
	roomFacility, err := models.GetAllRoomFacility()
	return roomFacility, err
}
