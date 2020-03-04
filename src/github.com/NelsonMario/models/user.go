package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
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
	Language string
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
_, err = GetApiKeyDetail(middleware.ApiKey)

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
	_, err = GetApiKeyDetail(middleware.ApiKey)
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
	_, err = GetApiKeyDetail(middleware.ApiKey)
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
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user := &User{FirstName: firstName, LastName: lastName, PhoneNumber:phoneNumber, Password:password, Email:email}
	db.Save(user)

	return user, nil
}

func UpdateUser(id int, firstName string, lastName string, phoneNumber string, email string, cityName string, address string, postcode string, language string) (*User, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user User
	db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{"first_name": firstName, "last_name": lastName, "phone_number": phoneNumber, "email": email, "city_name": cityName, "address": address, "post_code": postcode, "language": language})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&user)
	return &user, nil
}
