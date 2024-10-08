package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"github.com/jinzhu/gorm"
	"time"
)

type Admin struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"type:varchar(100);not null"`
	Email     string     `gorm:"type:varchar(100);not null"`
	Password  string     `gorm:"type:varchar(100);not null"`
}

var db *gorm.DB
var err error

func init() {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Admin{})
}

func GetAllAdmin() ([]Admin, error) {
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var admins []Admin
	db.Find(&admins)
	fmt.Println(admins)
	return admins, err
}

func GetAdmin(id int) ([]Admin, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		panic("Failed to connect to database")
	}

	defer db.Close()

	var admin []Admin
	db.Where("id = ?", id).Find(&admin)
	return admin, err
}

func InsertAdmin(name string, email string, password string) (*Admin, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	admin := &Admin{Name: name, Email: email, Password: password}
	db.Save(admin)

	return admin, nil
}

func UpdateAdmin(id int, name string, email string, password string) (*Admin, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var admin Admin
	db.Model(&admin).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "email": email, "password": password})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&admin)
	return &admin, nil
}

func RemoveAdmin(id int) (*Admin, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	admin := Admin{ID: id}
	db.Where("id = ?", id).Find(&admin)
	return &admin, db.Delete(admin).Error
}
