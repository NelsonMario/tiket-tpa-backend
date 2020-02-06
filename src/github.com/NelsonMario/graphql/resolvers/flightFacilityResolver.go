package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllFlightFacility(p graphql.ResolveParams) (i interface{}, e error) {
	facility, err := models.GetAllFacility()
	return facility, err
}
