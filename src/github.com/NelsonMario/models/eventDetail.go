package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type EventDetail struct{
	Id			int			`gorm:primary_key`
	EventRefer int
	Name string `gorm:"type:varchar(100)"`
	Price int
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
	db.AutoMigrate(&EventDetail{}).AddForeignKey("event_refer", "events(id)", "CASCADE", "CASCADE")
}

func InsertEventDetail(name string, eventRefer int, price int)(*EventDetail, error){
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	EventDetail := &EventDetail{Name:name, EventRefer: eventRefer, Price:price}

	db.Save(EventDetail)
	return EventDetail, err
}

func UpdateEventDetail(id int, name string, eventRefer int, price int) (*EventDetail, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var EventDetail EventDetail
	db.Model(&EventDetail).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "event_refer": eventRefer, "price": price})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&EventDetail)
	return &EventDetail, nil
}

func RemoveEventDetail(id int) (*EventDetail, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()
	fmt.Println(id)
	EventDetail := EventDetail{Id: id}
	db.Where("id = ?", id).Delete(&EventDetail)
	return &EventDetail, db.Delete(EventDetail).Error
}



