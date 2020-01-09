package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type User struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	FirstName      string     `gorm:"type:varchar(100);not null"`
	LastName     string     `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);not null"`
	Password  string     `gorm:"type:varchar(100);not null"`
	PhoneNumber  string     `gorm:"type:varchar(100);not null"`
	CityName string `gorm:"type:varchar(100);"`
	Address string `gorm:"type:varchar(100);"`
	PostCode string `gorm:"type:varchar(100);"`
}

func init() {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&User{})
}

func GetAllUser() ([]User, error) {
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Users []User
	db.Find(&Users)
	fmt.Println(Users)
	return Users, err
}
func GetUserById(id int) ([]User, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		panic("Failed to connect to database")
	}

	defer db.Close()

	var User []User
	db.Where("id = ?", id).Find(&User)
	return User, err
}

func GetUserByEmailOrPhone(input string) ([]User, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		panic("Failed to connect to database")
	}

	defer db.Close()

	var User []User
	db.Where("email = ?", input).Or("phone_number = ?", input).Find(&User)
	return User, err
}

func InsertUser(firstName string, lastName string, phoneNumber string, password string, email string) (*User, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	user := &User{FirstName: firstName, LastName: lastName, PhoneNumber:phoneNumber, Password:password, Email:email}
	db.Save(user)

	return user, nil
}
