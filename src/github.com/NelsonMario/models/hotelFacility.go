package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type HotelFacility struct{

	ID int `gorm: "primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Facility Facility `gorm:"foreignkey:FacilityRefer"`
	FacilityRefer uint
	HotelReferFacility uint
}

func init() {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&HotelFacility{})
}

func GetAllHotelFacility() ([]HotelFacility, error){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var HotelFacility []HotelFacility

	db.Find(&HotelFacility)
	for i,_ := range HotelFacility{
		db.Model(HotelFacility[i]).Related(&HotelFacility[i].Facility, "FacilityRefer")
	}

	fmt.Println(HotelFacility)
	return HotelFacility, err
}

func InsertHotelFacility(facilityRefer uint, hotelFacilityRefer uint)(*HotelFacility, error){
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	hotelFacility := &HotelFacility{FacilityRefer:facilityRefer, HotelReferFacility:hotelFacilityRefer}
	db.Save(hotelFacility)

	return hotelFacility, nil
}

func UpdateHotelFacility(Id int, facilityRefer uint, hotelFacilityRefer uint)(*HotelFacility, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotelFacility HotelFacility
	db.Model(&hotelFacility).Where("id = ?", Id).Updates(map[string]interface{}{"facility_refer": facilityRefer, "hotel_facility_refer":hotelFacilityRefer})
	db.Where("id = ?", Id).Find(&hotelFacility)
	return &hotelFacility, nil
}

func RemoveHotelFacility(id int) (*HotelFacility, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	hotelFacility := HotelFacility{ID: id}
	db.Where("id = ?", id).Find(&hotelFacility)
	return &hotelFacility, db.Delete(hotelFacility).Error
}


