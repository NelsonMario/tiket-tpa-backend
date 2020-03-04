package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Station struct{
	Id			int			`gorm:primary_key`
	Code		string		`gorm: "type:char(3);not null"`
	Name		string		`gorm: "type:varchar(100);not null"`
	City		string		`gorm: "type:varchar(100);not null"`
	CityCode	string		`gorm: "type:char(4);not null"`
	Province	string		`gorm: "type:varchar(100);not null"`
	Country		string		`gorm: "type:varchar(100);not null"`
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
	db.AutoMigrate(&Station{})
}

func GetAllStation() ([]Station, error) {
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Station []Station
	db.Find(&Station)
	fmt.Println(Station)
	return Station, err
}

func GetDistinctStation() ([]Station, error) {
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Station []Station
	db.Select("DISTINCT(city)").Find(&Station)
	fmt.Println(Station)
	return Station, err
}

func InsertStation(code string, name string, city string, cityCode string, province string, country string)(*Station, error){
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	station := &Station{Code:code, Name:name, City:city, CityCode:cityCode, Province:province, Country:country}
	db.Save(station)

	return station, nil
}