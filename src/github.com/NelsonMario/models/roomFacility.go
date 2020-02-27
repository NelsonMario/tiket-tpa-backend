package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type RoomFacility struct{

	ID int `gorm: "primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Facility Facility `gorm:"foreignkey:FacilityRefer"`
	FacilityRefer uint
	FacilityType string `gorm:"type:varchar(100);not null"`
	RoomReferFacility uint
}

func init() {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&RoomFacility{})
}

func GetAllRoomFacility() ([]RoomFacility, error){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var RoomFacility []RoomFacility

	db.Find(&RoomFacility)
	for i,_ := range RoomFacility{
		db.Model(RoomFacility[i]).Related(&RoomFacility[i].Facility, "FacilityRefer")
	}

	fmt.Println(RoomFacility)
	return RoomFacility, err
}


