package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type HotelRoom struct{

	ID int `gorm: "primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Room Room `gorm:"foreignkey:RoomRefer"`
	RoomRefer uint
	HotelReferRoom uint
}

func init() {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&HotelRoom{})
}

func GetAllHotelRoom() ([]HotelRoom, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var HotelRoom []HotelRoom

	db.Find(&HotelRoom)
	for i,_ := range HotelRoom{
		db.Model(HotelRoom[i]).Related(&HotelRoom[i].Room, "RoomRefer")
	}

	fmt.Println(HotelRoom)
	return HotelRoom, err
}



