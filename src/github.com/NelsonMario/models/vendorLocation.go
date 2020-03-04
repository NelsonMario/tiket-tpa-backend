package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type VendorLocation struct{
	Id int `gorm:"primary_key"`
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time	`sql:index`
	Location Location `gorm:"foreignkey:LocationRefer"`
	LocationRefer uint
	VendorLocationRefer uint
}

func init(){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&VendorLocation{})
}

func GetAllVendorLocation() ([]VendorLocation, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if(err != nil){
		panic(err)
	}

	var VendorLocation []VendorLocation

	db.Find(&VendorLocation)
	return VendorLocation, err
}