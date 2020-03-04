package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Review struct{
	Id			int			`gorm:primary_key`
	Description string
	Review int
	User string
	Category string
	Caption string
	HotelRefer int
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time	`sql:index`
}

func init() {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Review{})
}

func GetAllReview(id int)([]Review, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if(err != nil){
		panic(err)
	}

	var Review []Review

	db.Where("hotel_refer = ?", id).Find(&Review)
	return Review, err
}