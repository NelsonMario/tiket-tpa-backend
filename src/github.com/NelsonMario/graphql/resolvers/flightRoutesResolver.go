package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllFlightRoutes(p graphql.ResolveParams) (i interface{}, e error) {
	routes, err := models.GetAllRoute()
	return routes, err
}
