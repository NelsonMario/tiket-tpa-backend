package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Vendor struct{
	Id int `gorm:primary_key`
	Name  string `gorm:varchar(100) ; not null`
	VendorLocation []VendorLocation `gorm:"foreignkey:VendorLocationRefer"`
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
	db.AutoMigrate(&Vendor{})
}

func GetAllVendor() ([]Vendor, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if(err != nil){
		panic(err)
	}

	var Vendor []Vendor

	db.Find(&Vendor)
	for i,_ := range Vendor{
		db.Model(Vendor[i]).Related(&Vendor[i].VendorLocation, "VendorLocationRefer")
		for j,_ := range Vendor[i].VendorLocation{
			db.Model(Vendor[i].VendorLocation[j]).Related(&Vendor[i].VendorLocation[j].Location, "LocationRefer")
		}
	}
	return Vendor, err
}