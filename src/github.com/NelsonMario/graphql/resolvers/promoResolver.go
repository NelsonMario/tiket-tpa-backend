package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllPromo(p graphql.ResolveParams)(i interface{}, e error){
	promo, err :=  models.GetAllPromo()
	return promo, err
}

func GetPromoById(p graphql.ResolveParams)(i interface{}, e error){
	id := p.Args["id"].(int)
	promo, err :=  models.GetPromoById(id)
	return promo, err
}