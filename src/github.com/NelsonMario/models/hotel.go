package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Hotel struct{
	Id			int			`gorm:primary_key`
	Name		string		`gorm: "type:varchar(100);not null"`
	Rating		int		`gorm: "type:int(100);not null"`
	Type string `gorm: "type:varchar(100);not null"`
	HotelFacility	[]HotelFacility `gorm:"foreignkey:HotelReferFacility"`
	HotelRoom []HotelRoom `gorm:"foreignkey:HotelReferRoom"`
	Location Location
	LocationRefer int
	HotelLat float64
	HotelLng float64
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time	`sql:index`
}

func init(){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Hotel{})
}

func GetAllHotel() ([]Hotel, error) {
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Hotel []Hotel
	db.Find(&Hotel)
	for i,_ := range Hotel{
		db.Model(Hotel[i]).Related(&Hotel[i].HotelFacility, "HotelReferFacility").Related(&Hotel[i].HotelRoom, "HotelReferRoom").Related(&Hotel[i].Location, "LocationRefer")
		for j,_ := range Hotel[i].HotelFacility{
			db.Model(Hotel[i].HotelFacility[j]).Related(&Hotel[i].HotelFacility[j].Facility, "FacilityRefer")
		}
		for j,_ := range Hotel[i].HotelRoom{
			db.Model(Hotel[i].HotelRoom[j]).Related(&Hotel[i].HotelRoom[j].Room, "RoomRefer")
			db.Model(Hotel[i].HotelRoom[j].Room).Related(&Hotel[i].HotelRoom[j].Room.RoomFacility, "RoomReferFacility")
			for k,_ := range Hotel[i].HotelRoom[j].Room.RoomFacility{
				db.Model(Hotel[i].HotelRoom[j].Room.RoomFacility[k]).Related(&Hotel[i].HotelRoom[j].Room.RoomFacility[k].Facility, "FacilityRefer")
			}
		}
	}
	fmt.Println(Hotel)
	return Hotel, err
}

func GetHotelById(id int) ([]Hotel, error) {
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Hotel []Hotel
	db.Where("id = ?", id).Find(&Hotel)
	for i,_ := range Hotel{
		db.Model(Hotel[i]).Related(&Hotel[i].HotelFacility, "HotelReferFacility").Related(&Hotel[i].HotelRoom, "HotelReferRoom").Related(&Hotel[i].Location, "LocationRefer")
		for j,_ := range Hotel[i].HotelFacility{
			db.Model(Hotel[i].HotelFacility[j]).Related(&Hotel[i].HotelFacility[j].Facility, "FacilityRefer")
		}
		for j,_ := range Hotel[i].HotelRoom{
			db.Model(Hotel[i].HotelRoom[j]).Related(&Hotel[i].HotelRoom[j].Room, "RoomRefer")
			db.Model(Hotel[i].HotelRoom[j].Room).Related(&Hotel[i].HotelRoom[j].Room.RoomFacility, "RoomReferFacility")
			for k,_ := range Hotel[i].HotelRoom[j].Room.RoomFacility{
				db.Model(Hotel[i].HotelRoom[j].Room.RoomFacility[k]).Related(&Hotel[i].HotelRoom[j].Room.RoomFacility[k].Facility, "FacilityRefer")
			}
		}
	}
	fmt.Println(Hotel)
	return Hotel, err
}


func GetDistinctHotel() ([]Hotel, error) {
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Hotel []Hotel
	db.Select("DISTINCT(city)").Find(&Hotel)
	return Hotel, err
}

func GetNearestHotelByLocation(location string) ([]Hotel, error) {
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Location Location
	db.Where("city = ?", location).Find(&Location)

	var Hotel []Hotel
	db.Where("25 > (6371 * acos(cos(radians(?)) * cos(radians(hotel_lat)) * cos(radians(hotel_lng) - radians(?)) +sin(radians(?)) * sin(radians(hotel_lat))))", Location.Lat, Location.Lng, Location.Lat).Find(&Hotel)
	for i,_ := range Hotel{
		db.Model(Hotel[i]).Related(&Hotel[i].HotelFacility, "HotelReferFacility").Related(&Hotel[i].HotelRoom, "HotelReferRoom").Related(&Hotel[i].Location, "LocationRefer")
		for j,_ := range Hotel[i].HotelFacility{
			db.Model(Hotel[i].HotelFacility[j]).Related(&Hotel[i].HotelFacility[j].Facility, "FacilityRefer")
		}
		for j,_ := range Hotel[i].HotelRoom{
			db.Model(Hotel[i].HotelRoom[j]).Related(&Hotel[i].HotelRoom[j].Room, "RoomRefer")
			db.Model(Hotel[i].HotelRoom[j].Room).Related(&Hotel[i].HotelRoom[j].Room.RoomFacility, "RoomReferFacility")
			for k,_ := range Hotel[i].HotelRoom[j].Room.RoomFacility{
				db.Model(Hotel[i].HotelRoom[j].Room.RoomFacility[k]).Related(&Hotel[i].HotelRoom[j].Room.RoomFacility[k].Facility, "FacilityRefer")
			}
		}
	}
	return Hotel, err
}

func GetHotelByLocation(location string) ([]Hotel, error) {
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Location Location
	db.Where("city = ?", location).Find(&Location)

	var Hotel []Hotel
	db.Where("id = ?", Location.Id).Find(&Hotel)
	for i,_ := range Hotel{
		db.Model(Hotel[i]).Related(&Hotel[i].HotelFacility, "HotelReferFacility").Related(&Hotel[i].HotelRoom, "HotelReferRoom").Related(&Hotel[i].Location, "LocationRefer")
		for j,_ := range Hotel[i].HotelFacility{
			db.Model(Hotel[i].HotelFacility[j]).Related(&Hotel[i].HotelFacility[j].Facility, "FacilityRefer")
		}
		for j,_ := range Hotel[i].HotelRoom{
			db.Model(Hotel[i].HotelRoom[j]).Related(&Hotel[i].HotelRoom[j].Room, "RoomRefer")
			db.Model(Hotel[i].HotelRoom[j].Room).Related(&Hotel[i].HotelRoom[j].Room.RoomFacility, "RoomReferFacility")
			for k,_ := range Hotel[i].HotelRoom[j].Room.RoomFacility{
				db.Model(Hotel[i].HotelRoom[j].Room.RoomFacility[k]).Related(&Hotel[i].HotelRoom[j].Room.RoomFacility[k].Facility, "FacilityRefer")
			}
		}
	}
	return Hotel, err
}

func InsertHotel(name string, rating int, _type string, locationRefer int, lat float64, lng float64)(*Hotel, error){
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	hotel := &Hotel{Name:name, Rating:rating, Type:_type, LocationRefer:locationRefer,HotelLat: lat, HotelLng: lng}
	db.Save(hotel)

	return hotel, nil
}

func UpdateHotel(Id int, name string, rating int)(*Hotel, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var Hotel Hotel
	db.Model(&Hotel).Where("id = ?", Id).Updates(map[string]interface{}{"name": name, "rating":rating})
	db.Where("id = ?", Id).Find(&Hotel)
	return &Hotel, nil
}

func RemoveHotel(id int) (*Hotel, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	hotel := Hotel{Id: id}
	db.Where("id = ?", id).Delete(&hotel)
	return &hotel, db.Delete(hotel).Error
}
