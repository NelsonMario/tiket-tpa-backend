package models

import (
	"github.com/NelsonMario/connection"
	"time"
)

type Airline struct{

	ID int `gorm: "primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name string `gorm: "type: varchar(100);not null"`
}

func init(){
	db, err = connection.ConnectDatabase();

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Airline{})
}

func GetAllAirline() ([]Airline, error){
	db, err = connection.ConnectDatabase()

	if(err != nil){
		panic(err)
	}

	var Airline []Airline

	db.Find(&Airline)
	return Airline, err
}

func GetAirlineById(id int)([]Airline, error){
	db, err = connection.ConnectDatabase()

	if(err != nil){
		panic(err)
	}

	defer db.Close()

	var Airline[] Airline

	db.Where("id = ?", id).Find(&Airline)
	return Airline, err
}

func InsertAirline(name string)(*Airline, error){
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	airline := &Airline{Name:name}

	db.Save(airline)
	return airline, err
}

