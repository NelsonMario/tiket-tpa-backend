package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type FlightFacility struct{

	ID int `gorm: "primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Facility Facility `gorm:"foreignkey:FacilityRefer"`
	FacilityRefer uint
	FlightReferFacility uint
}

func init() {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&FlightFacility{})
}

func GetAllFlightFacility() ([]FlightFacility, error){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var FlightFacility []FlightFacility

	db.Find(&FlightFacility)
	for i,_ := range FlightFacility{
		db.Model(FlightFacility[i]).Related(&FlightFacility[i].Facility, "FacilityRefer")
	}

	fmt.Println(FlightFacility)
	return FlightFacility, err
}

func InsertFlightFacility(facilityRefer uint, flightFacilityRefer uint)(*FlightFacility, error){
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	flightFacility := &FlightFacility{FacilityRefer: facilityRefer, FlightReferFacility:flightFacilityRefer}

	db.Save(flightFacility)
	return flightFacility, err
}

func UpdateFlightFacility(id int, facilityRefer uint, flightFacilityRefer uint) (*FlightFacility, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var flightFacility FlightFacility
	db.Model(&flightFacility).Where("id = ?", id).Updates(map[string]interface{}{"facility_refer": facilityRefer, "flight_facility_refer": flightFacilityRefer})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&flightFacility)
	return &flightFacility, nil
}

func RemoveFlightFacility(id int) (*FlightFacility, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	flightFacility := FlightFacility{ID: id}
	db.Where("id = ?", id).Find(&flightFacility)
	return &flightFacility, db.Delete(flightFacility).Error
}