package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Train struct{

	ID int `gorm: "primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name string `gorm: "type: varchar(100);not null"`
}

func init(){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Train{})
}

func GetAllTrain() ([]Train, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if(err != nil){
		panic(err)
	}

	var Train []Train

	db.Find(&Train)
	return Train, err
}

func GetTrainById(id int)([]Train, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if(err != nil){
		panic(err)
	}

	defer db.Close()

	var Train[] Train

	db.Where("id = ?", id).Find(&Train)
	return Train, err
}

func InsertTrain(name string)(*Train, error){
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	Train := &Train{Name:name}

	db.Save(Train)
	return Train, err
}

