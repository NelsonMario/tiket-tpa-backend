package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Bank struct{

	ID int `gorm: "primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name string `gorm: "type: varchar(100);not null"`
	Desc string
}

func init(){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Bank{})
}
func GetAllBanks() ([]Bank, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil ,err
	}

	defer db.Close()

	var banks []Bank

	db.Find(&banks)

	return banks, nil
}

