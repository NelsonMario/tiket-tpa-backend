package resolvers

import (
	"fmt"
	"github.com/NelsonMario/models"
	"github.com/graphql-go/graphql"
)

func GetSingleBlog(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	blog, err := models.GetSingleBlog(id)
	return blog, err
}

func GetAllBlog(p graphql.ResolveParams) (i interface{}, e error) {
	blog, err := models.GetAllBlog()
	return blog, err
}

func InsertBlog(p graphql.ResolveParams)(i interface{}, e error){

	userId := p.Args["userId"].(int)
	title := p.Args["title"].(string)
	value := p.Args["value"].(string)
	image := p.Args["image"].(string)
	thumbnail := p.Args["thumbnail"].(string)
	viewer := p.Args["viewer"].(int)

	if(userId == -1 || title == " " || value == " " || image == " " || thumbnail == " "){
		fmt.Println("ASDs")
		return 1, nil
	}

	airport, err := models.InsertBlog(userId, title, value, image, thumbnail, viewer)

	if err != nil {
		return nil, err
	}
	return airport, nil
}

func UpdateBlog(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	title := p.Args["title"].(string)
	thumbnail := p.Args["thumbnail"].(string)
	value := p.Args["value"].(string)


	blog, err := models.UpdateBlog(id, title, thumbnail, value)

	if err != nil {
		return nil, err
	}
	return blog, nil
}


func RemoveBlog(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)

	blog, err := models.RemoveBlog(id)

	if err != nil {
		return nil, err
	}
	return blog, nil
}

func IncreaseViewer(p graphql.ResolveParams)(i interface{}, e error){
	id := p.Args["id"].(int)
	viewer := p.Args["viewer"].(int)
	blog, err := models.IncreaseViewer(id, viewer)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

