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

func InsertFacility(name string)(*Facility, error){
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	facility := &Facility{Name:name}

	db.Save(facility)
	return facility, err
}

func UpdateFacility(id int, name string) (*Facility, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var facility Facility
	db.Model(&facility).Where("id = ?", id).Updates(map[string]interface{}{"name": name})
	db.Where("id = ?", id).Find(&facility)
	return &facility, nil
}

func RemoveFacility(id int) (*Facility, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	facility := Facility{ID: id}
	db.Where("id = ?", id).Find(&facility)
	return &facility, db.Delete(facility).Error
}

