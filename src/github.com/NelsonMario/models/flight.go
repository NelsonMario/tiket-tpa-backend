package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Flight struct {
	Id				int			`gorm:primary_key`
	Airline 		Airline		`gorm:"foreignkey:AirlineRefer"`
	AirlineRefer	int
	From			Airport		`gorm:"foreignkey:FromRefer"`
	FromRefer		int
	To				Airport		`gorm:"foreignkey:ToRefer"`
	ToRefer			int
	Departure		time.Time
	Arrival 		time.Time
	Duration		int
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

func InsertFlight(airlineRefer int, fromRefer int, toRefer int, departure time.Time, arrival time.Time, duration int, price int, tax int, serviceCharge int)(*Flight, error){
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	flight := &Flight{AirlineRefer: airlineRefer, FromRefer: fromRefer, ToRefer: toRefer, Departure: departure, Arrival: arrival, Duration: duration, Price: price, Tax: tax, ServiceCharge: serviceCharge}
	db.Save(flight)

	return flight, nil
}

func GetAllFlight() ([]Flight, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

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

	return Flights, err
}

func GetFlightByFromAndTo(from string, to string) ([]Flight, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

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

func GetFlightByOneSchedule(from string, to string, date string) ([]Flight, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Flights []Flight

	db.Where("from_refer IN (?) AND to_refer IN (?) AND date_trunc('day', departure) = (?)", db.Table("airports").Select("id").Where("city = ?", from).SubQuery(), db.Table("airports").Select("id").Where("city = ?", to).SubQuery(), date).Find(&Flights)

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

func GetFlightBySchedule(fromCity string, toCity string, fromDate string, toDate string) ([]Flight, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)


	if toDate == ""{
		toDate = time.Now().Format("2006-01-02")
	}

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Flights []Flight
	db.Where("((to_refer IN (?) AND from_refer IN (?)) OR (to_refer IN (?) AND from_refer IN (?))) AND ((date_trunc('day', departure) = (?)) OR (date_trunc('day', departure) = (?)))", db.Table("airports").Select("id").Where("city = ?", toCity).SubQuery(), db.Table("airports").Select("id").Where("city = ?", fromCity).SubQuery(), db.Table("airports").Select("id").Where("city = ?", fromCity).SubQuery(), db.Table("airports").Select("id").Where("city = ?", toCity).SubQuery(), fromDate ,toDate).Find(&Flights)
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

func UpdateFlight(Id int, airlineRefer int, fromRefer int, toRefer int, departure time.Time, arrival time.Time, duration int, price int, tax int, serviceCharge int)(*Flight, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
	return nil, err
	}
	defer db.Close()

	var flight Flight
	db.Model(&flight).Where("id = ?", Id).Updates(map[string]interface{}{"airline_refer": airlineRefer, "from_refer":fromRefer, "to_refer":toRefer, "departure": departure, "arrival": arrival, "duration": duration, "price":price, "tax":tax, "service_charge": serviceCharge})
	db.Where("id = ?", Id).Find(&flight)
	return &flight, nil
}

func RemoveFlight(id int) (*Flight, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	fmt.Println(id)
	flight := Flight{Id: id}
	db.Where("id = ?", id).Delete(&flight)
	return &flight, db.Delete(flight).Error
}
