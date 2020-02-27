package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllHotelFacility(p graphql.ResolveParams) (i interface{}, e error) {
	hotelFacility, err := models.GetAllHotelFacility()
	return hotelFacility, err
}
