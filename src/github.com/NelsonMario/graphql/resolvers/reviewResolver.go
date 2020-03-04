package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllReview(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	review, err := models.GetAllReview(id)
	return review, err
}

