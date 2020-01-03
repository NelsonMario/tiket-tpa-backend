package resolvers

import (
	 "github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllAdmin(p graphql.ResolveParams) (i interface{}, e error) {
	admins, err := models.GetAllAdmin()
	return admins, err
}

func GetAdmin(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	admin, err := models.GetAdmin(id)
	return admin, err
}

func InsertAdmin(p graphql.ResolveParams) (i interface{}, e error) {
	name := p.Args["name"].(string)
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)

	admin, err := models.InsertAdmin(name, email, password)

	if err != nil {
		return nil, err
	}
	return admin, nil
}

func UpdateAdmin(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	name := p.Args["name"].(string)
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)

	admin, err := models.UpdateAdmin(id, name, email, password)

	if err != nil {
		return nil, err
	}
	return admin, nil
}

func RemoveAdmin(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)

	admin, err := models.RemoveAdmin(id)

	if err != nil {
		return nil, err
	}
	return admin, nil
}
