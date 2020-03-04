package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type VendorCar struct{
	Id int `gorm:primary_key`
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time	`sql:index`
	Vendor Vendor `gorm:"foreignkey:VendorRefer"`
	VendorRefer uint
	VendorCarRefer uint
}

func init(){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&VendorCar{})
}

func GetAllVendorCar() ([]VendorCar, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if(err != nil){
		panic(err)
	}

	var VendorCar []VendorCar

	db.Find(&VendorCar)
	return VendorCar, err
}