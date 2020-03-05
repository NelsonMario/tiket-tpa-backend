package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllOrder(p graphql.ResolveParams)(i interface{}, e error){
	order, err :=  models.GetAllOrder()
	return order, err
}

func GetOrderById(p graphql.ResolveParams)(i interface{}, e error){
	id := p.Args["id"].(int)
	order, err :=  models.GetOrderById(id)
	return order, err
}