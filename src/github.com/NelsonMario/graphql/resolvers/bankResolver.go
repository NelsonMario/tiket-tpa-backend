package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllBank(p graphql.ResolveParams)(i interface{},  e error){
	bank, err := models.GetAllBanks()
	return bank, err
}
