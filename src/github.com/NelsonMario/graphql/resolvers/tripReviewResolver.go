package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllTripReview(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	review, err := models.GetAllTripReview(id)
	return review, err
}

