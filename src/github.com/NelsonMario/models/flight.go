package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type Flight struct {
	Id				int			`gorm:primary_key`
	Airline 		Airline		`gorm:"foreignkey:AirlineRefer"`
	AirlineRefer	uint
	From			Airport		`gorm:"foreignkey:FromRefer"`
	FromRefer		uint
	To				Airport		`gorm:"foreignkey:ToRefer"`
	ToRefer			uint
	Departure		time.Time
	Arrival 		time.Time
	Duration		uint
	Price			int				`gorm:"type:int"`
	Tax				int				`gorm:"type:int"`
	ServiceCharge	int				`gorm:"type:int"`
	FlightRoutes	[]FlightRoutes `gorm:"foreignkey:FlightReferRoute"`
	FlightFacility	[]FlightFacility `gorm:"foreignkey:FlightReferFacility"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time		`sql:index`

}

func init(){

	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Flight{}).AddForeignKey("airline_refer", "airlines(id)", "CASCADE", "CASCADE").AddForeignKey("from_refer", "airports(id)", "CASCADE", "CASCADE").AddForeignKey("to_refer", "airports(id)", "CASCADE", "CASCADE")
}

func InsertFlight(airlineRefer uint, fromRefer uint, toRefer uint, depature time.Time, arrival time.Time, duration uint, price int, tax int, serviceCharge int)(*Flight, error){
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	flight := &Flight{AirlineRefer: airlineRefer, FromRefer: fromRefer, ToRefer: toRefer, Departure: depature, Arrival: arrival, Duration: duration, Price: price, Tax: tax, ServiceCharge: serviceCharge}
	db.Save(flight)

	return flight, nil
}

func GetAllFlight() ([]Flight, error){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Flights []Flight

	db.Find(&Flights)
	for i,_ := range Flights{
		db.Model(Flights[i]).Related(&Flights[i].Airline, "AirlineRefer").Related(&Flights[i].From, "FromRefer").Related(&Flights[i].To, "ToRefer").Related(&Flights[i].FlightRoutes, "FlightReferRoute").Related(&Flights[i].FlightFacility, "FlightReferFacility")
		for j,_ := range Flights[i].FlightRoutes{
			db.Model(Flights[i].FlightRoutes[j]).Related(&Flights[i].FlightRoutes[j].Airport, "AirportRefer")
		}
		for j,_ := range Flights[i].FlightFacility{
			db.Model(Flights[i].FlightFacility[j]).Related(&Flights[i].FlightFacility[j].Facility, "FacilityRefer")
		}
	}

	fmt.Println(Flights)
	return Flights, err
}

func GetFlightByFromAndTo(from string, to string) ([]Flight, error){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Flights []Flight

	db.Where("from_refer IN (?) AND to_refer IN (?)", db.Table("airports").Select("id").Where("city = ?", from).SubQuery(), db.Table("airports").Select("id").Where("city = ?", to).SubQuery()).Find(&Flights)
	for i,_ := range Flights{
		db.Model(Flights[i]).Related(&Flights[i].Airline, "AirlineRefer").Related(&Flights[i].From, "FromRefer").Related(&Flights[i].To, "ToRefer").Related(&Flights[i].FlightRoutes, "FlightReferRoute").Related(&Flights[i].FlightFacility, "FlightReferFacility")
		for j,_ := range Flights[i].FlightRoutes{
			db.Model(Flights[i].FlightRoutes[j]).Related(&Flights[i].FlightRoutes[j].Airport, "AirportRefer")
		}
		for j,_ := range Flights[i].FlightFacility{
			db.Model(Flights[i].FlightFacility[j]).Related(&Flights[i].FlightFacility[j].Facility, "FacilityRefer")
		}
	}
	return Flights, err
}


