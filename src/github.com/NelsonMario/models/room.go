package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Room struct {
	Id				int			`gorm:primary_key`
	CheckIn			time.Time
	CheckOut 		time.Time
	Bed 			string	`gorm:"type:varchar(100); not null"`
	Price			int				`gorm:"type:int"`
	MaxGuest 		int 			`gorm:"type:int"`
	Size			int				`gorm:"type:int"`
	RoomFacility	[]RoomFacility `gorm:"foreignkey:RoomReferFacility"`
	Type string
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time		`sql:index`

}

func init(){

	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Room{})
}

func InsertRoom(checkin time.Time, checkout time.Time, bed string, price int, maxGuest int, size int)(*Room, error){
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	room := &Room{CheckIn:checkin, CheckOut:checkout, Bed:bed, Price:price, MaxGuest:maxGuest, Size:size}
	db.Save(room)

	return room, nil
}

func GetAllRoom() ([]Room, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Room []Room

	db.Find(&Room)
	for i,_ := range Room{
		db.Model(Room[i]).Related(&Room[i].RoomFacility, "RoomReferFacility")
		for j,_ := range Room[i].RoomFacility{
			db.Model(Room[i].RoomFacility[j]).Related(&Room[i].RoomFacility[j].Facility, "FacilityRefer")
		}
	}

	return Room, err
}

func GetRoomBySchedule(location int, fromDate string, toDate string) ([]Room, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)


	if toDate == ""{
		toDate = time.Now().Format("2006-01-02")
	}

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Room []Room
	db.Where("hotel_refer = (?) AND ((date_trunc('day', departure) = (?)) OR (date_trunc('day', departure) = (?)))",  db.Table("hotels").Select("id").Where("City = ?", location).SubQuery(), fromDate ,toDate).Find(&Room)
	for i,_ := range Room{
		db.Model(Room[i]).Related(&Room[i].RoomFacility, "RoomReferFacility")
		for j,_ := range Room[i].RoomFacility{
			db.Model(Room[i].RoomFacility[j]).Related(&Room[i].RoomFacility[j].Facility, "FacilityRefer")
		}
	}
	return Room, err
}


