package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type Facility struct {
	ID        int `gorm:"primary_key"`
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
	db.AutoMigrate(&Facility{})
}

func GetAllFacility() ([]Facility, error) {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Facility []Facility
	db.Find(&Facility)
	fmt.Println(Facility)
	return Facility, err
}
