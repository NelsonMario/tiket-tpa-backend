package resolvers

import (
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetAllUser(p graphql.ResolveParams) (i interface{}, e error) {
	users, err := models.GetAllUser()
	return users, err
}

//func GetUser(p graphql.ResolveParams) (i interface{}, e error) {
//	id := p.Args["id"].(int)
//	user, err := models.GetAdmin(id)
//	return user, err
//}

func InsertUser(p graphql.ResolveParams) (i interface{}, e error) {
	firstName := p.Args["firstName"].(string)
	lastName := p.Args["lastName"].(string)
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)
	phoneNumber := p.Args["phoneNumber"].(string)


	user, err := models.InsertUser(firstName, lastName, phoneNumber, password, email)

	if err != nil {
		return nil, err
	}
	return user, nil
}
