package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type TripReview struct{
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
	db.AutoMigrate(&TripReview{})
}

func GetAllTripReview(id int)([]TripReview, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if(err != nil){
		panic(err)
	}

	var TripReview []TripReview

	db.Where("id = ?", id).Find(&TripReview)
	return TripReview, err
}