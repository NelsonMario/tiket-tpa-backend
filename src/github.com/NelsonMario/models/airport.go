package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type Airport struct{
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
	db.AutoMigrate(&Airport{})
}

func GetAllAirport() ([]Airport, error) {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Airport []Airport
	db.Find(&Airport)
	fmt.Println(Airport)
	return Airport, err
}

func GetDistinctAirport() ([]Airport, error) {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Airport []Airport
	db.Select("DISTINCT(city)").Find(&Airport)
	fmt.Println(Airport)
	return Airport, err
}

func InsertAirport(code string, name string, city string, cityCode string, province string, country string)(*Airport, error){
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	airport := &Airport{Code:code, Name:name, City:city, CityCode:cityCode, Province:province, Country:country}
	db.Save(airport)

	return airport, nil
}

func GetAirportById(id int)([]Airport, error){
	db, err = connection.ConnectDatabase()

	if(err != nil){
		panic(err)
	}

	defer db.Close()

	var Airport[] Airport

	db.Where("id = ?", id).Find(&Airport)
	return Airport, err
}

func UpdateAirport(id int, code string, name string, city string, cityCode string, province string, country string) (*Airport, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var airport Airport
	db.Model(&airport).Where("id = ?", id).Updates(map[string]interface{}{"code": code, "name": name, "city": city, "city_code": cityCode, "province": province, "country": country})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&airport)
	return &airport, nil
}

func RemoveAirport(id int) (*Airport, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	airport := Airport{Id: id}
	db.Where("id = ?", id).Find(&airport)
	return &airport, db.Delete(airport).Error
}