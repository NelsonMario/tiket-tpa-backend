package models

import (
	"github.com/NelsonMario/connection"
	"time"
)

type Order struct{
	Id			int			`gorm:primary_key`
	User User
	UserRefer int
	Name string
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time	`sql:index`
}

func init(){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Order{})
}