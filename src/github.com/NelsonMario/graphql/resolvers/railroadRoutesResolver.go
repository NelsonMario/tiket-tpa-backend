package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllRailroadRoutes(p graphql.ResolveParams) (i interface{}, e error) {
	routes, err := models.GetAllRailroadRoute()
	return routes, err
}
