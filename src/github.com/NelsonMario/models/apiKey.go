package models

import (
	"github.com/NelsonMario/connection"
	"time"
)

type ApiKey struct {
	ID int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Key string `gorm:"type:varchar(100);unique_index"`
}

func init(){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&ApiKey{})
}

func GetApiKeyDetail(key string) ([]ApiKey, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()
	var apiKey []ApiKey
	db.Where("key = ?", key).Find(&apiKey)

	if len(apiKey) == 1 {
		return apiKey, err
	} else {
		return nil, connection.ReturnError("Key not- valid")
	}
}

