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


