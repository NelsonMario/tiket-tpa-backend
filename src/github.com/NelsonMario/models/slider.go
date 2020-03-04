package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Slider struct {
	Id       int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"type:varchar(100);not null"`
}

func init() {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Slider{})
}

func GetAllSlider() ([]Slider, error) {
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var sliders []Slider
	db.Find(&sliders)
	fmt.Println(sliders)
	return sliders, err
}

func InsertSlider(name string) (*Slider, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	slider := &Slider{Name: name}
	db.Save(slider)

	return slider, nil
}
