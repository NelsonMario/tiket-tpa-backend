package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type FlightRoutes struct{

	ID int `gorm: "primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Airport Airport `gorm:"foreignkey:AirportRefer"`
	AirportRefer uint
	FlightReferRoute uint
}

func init(){
	db, err = connection.ConnectDatabase();

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&FlightRoutes{}).AddForeignKey("airport_refer", "airports(id)", "CASCADE", "CASCADE")
}

func GetAllRoute() ([]FlightRoutes, error){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var FlightRoutes []FlightRoutes

	db.Find(&FlightRoutes)
	for i,_ := range FlightRoutes{
		db.Model(FlightRoutes[i]).Related(&FlightRoutes[i].Airport, "AirportRefer")
	}

	fmt.Println(FlightRoutes)
	return FlightRoutes, err
}