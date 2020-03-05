package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Order struct{
	Id			int			`gorm:primary_key`
	UserRefer int
	HotelRefer int
	CheckIn time.Time
	CheckOut time.Time
	TotalGuest int
	ConfirmationCode string
	TotalNightFee int
	CleaningFee int
	ServiceFee int
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time `sql:index`
}

func init(){

	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Order{})
}

func GetAllOrder() ([]Order, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Order []Order

	db.Find(&Order)
	return Order, err
}

func GetOrderById(id int) ([]Order, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Order []Order

	db.Where("user_refer = ?", id).Find(&Order)

	return Order, err
}

