package models

import (
	"github.com/NelsonMario/connection"
	"time"
)

type Location struct{
	Id			int			`gorm:primary_key`
	City		string		`gorm: "type:varchar(100);not null"`
	Country		string		`gorm: "type:varchar(100);not null"`
	Lat float64
	Lng	float64
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
	db.AutoMigrate(&Location{})
}

func GetAllLocation() ([]Location, error){
	db, err = connection.ConnectDatabase()

	if(err != nil){
		panic(err)
	}

	var Location []Location

	db.Find(&Location)
	return Location, err
}

func InsertLocation(city string, country string, lat float64, lng float64)(*Location, error){
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	location := &Location{City:city, Country:country, Lat:lat, Lng:lng}

	db.Save(location)
	return location, err
}