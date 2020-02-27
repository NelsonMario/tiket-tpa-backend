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

func InsertFlightRoutes(airportRefer uint, flightReferRoutes uint)(*FlightRoutes, error){
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	FlightRoutes := &FlightRoutes{AirportRefer:airportRefer, FlightReferRoute:flightReferRoutes}

	db.Save(FlightRoutes)
	return FlightRoutes, err
}

func UpdateFlightRoutes(id int, airportRefer uint, flightReferRoutes uint) (*FlightRoutes, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var flightRoutes FlightRoutes
	db.Model(&flightRoutes).Where("id = ?", id).Updates(map[string]interface{}{"airport_refer": airportRefer, "flight_refer_airport": flightReferRoutes})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&flightRoutes)
	return &flightRoutes, nil
}

func RemoveFlightRoutes(id int) (*FlightRoutes, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	flightRoutes := FlightRoutes{ID: id}
	db.Where("id = ?", id).Find(&flightRoutes)
	return &flightRoutes, db.Delete(flightRoutes).Error
}