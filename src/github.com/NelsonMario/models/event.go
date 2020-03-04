package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Event struct{
	Id			int			`gorm:primary_key`
	Name string `gorm:"type:varchar(100)"`
	StartDate time.Time
	EndDate time.Time
	Location Location
	LocationRefer int
	EventLat float64
	EventLng float64
	EventDetail []EventDetail `gorm:"foreignkey:EventRefer"`
	Picture string `gorm:"type:varchar(255)"`
	Description string
	Category string `gorm:"type:varchar(50)"`
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time `sql:index`
}

func init(){

	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Event{})
}

func GetAllEvent() ([]Event, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Event []Event

	db.Find(&Event)
	for i,_ := range Event{
		db.Model(Event[i]).Related(&Event[i].Location, "LocationRefer").Related(&Event[i].EventDetail, "EventRefer")
	}

	return Event, err
}

func GetEventById(id int) ([]Event, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Event []Event

	db.Where("id = ?", id).Find(&Event)
	for i,_ := range Event{
		db.Model(Event[i]).Related(&Event[i].Location, "LocationRefer").Related(&Event[i].EventDetail, "EventRefer")
	}

	return Event, err
}

func GetNearestEventByLocation(location string) ([]Event, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Location Location
	db.Where("city = ?", location).Find(&Location)

	var Event []Event
	db.Where("150 > (6371 * acos(cos(radians(?)) * cos(radians(event_lat)) * cos(radians(event_lng) - radians(?)) +sin(radians(?)) * sin(radians(event_lat))))", Location.Lat, Location.Lng, Location.Lat).Find(&Event)
	for i,_ := range Event{
		db.Model(Event[i]).Related(&Event[i].Location, "LocationRefer").Related(&Event[i].EventDetail, "EventRefer")
	}
	return Event, err
}

func GetEventByLocation(location string) ([]Event, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Location Location
	db.Where("city = ?", location).Find(&Location)

	var Event []Event
	db.Where("location_refer = ?", Location.Id).Find(&Event)
	for i,_ := range Event{
		db.Model(Event[i]).Related(&Event[i].Location, "LocationRefer").Related(&Event[i].EventDetail, "EventRefer")
	}
	return Event, err
}

func GetAllEventByCategory(category string) ([]Event, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Event []Event

	db.Where("category = ?", category).Order("start_date").Limit(3).Find(&Event)
	for i,_ := range Event{
		db.Model(Event[i]).Related(&Event[i].Location, "LocationRefer").Related(&Event[i].EventDetail, "EventRefer")
	}

	return Event, err
}

func InsertEvent(name string, locationRefer int, startDate time.Time, endDate time.Time, eventLat float64, eventLng float64, category string, description string)(*Event, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	event := &Event{Name:name, LocationRefer: locationRefer, StartDate:startDate, EndDate:endDate, EventLat:eventLat, EventLng:eventLng, Category:category, Description:description}

	db.Save(event)
	return event, err
}

func UpdateEvent(id int, name string, locationRefer int, startDate time.Time, endDate time.Time) (*Event, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var event Event
	db.Model(&event).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "location_refer": locationRefer, "startDate": startDate, "endDate":endDate})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&event)
	return &event, nil
}

func RemoveEvent(id int) (*Event, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()
	fmt.Println(id)
	event := Event{Id: id}
	db.Where("id = ?", id).Delete(&event)
	return &event, db.Delete(event).Error
}
