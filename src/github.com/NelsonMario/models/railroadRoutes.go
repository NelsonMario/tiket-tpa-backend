package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type RailroadRoutes struct{

	ID int `gorm: "primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Station Station `gorm:"foreignkey:StationRefer"`
	StationRefer uint
	RailroadReferRoute uint
}

func init(){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&RailroadRoutes{}).AddForeignKey("station_refer", "stations(id)", "CASCADE", "CASCADE")
}

func GetAllRailroadRoute() ([]RailroadRoutes, error){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var railroadRoutes []RailroadRoutes

	db.Find(&railroadRoutes)
	for i,_ := range railroadRoutes{
		db.Model(railroadRoutes[i]).Related(&railroadRoutes[i].Station, "StationRefer")
	}

	fmt.Println(railroadRoutes)
	return railroadRoutes, err
}