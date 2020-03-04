package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Car struct{
	Id int `gorm:primary_key`
	Name  string `gorm:varchar(100) ; not null`
	Model string `gorm:varchar(100) ; not null`
	Passanger  int `gorm:int;`
	Price  int `gorm:int;`
	VendorCar []VendorCar `gorm:"foreignkey:VendorCarRefer"`
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
	db.AutoMigrate(&Car{})
}

func GetAllCar() ([]Car, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Car []Car

	db.Find(&Car)
	for i,_ := range Car{
		db.Model(Car[i]).Related(&Car[i].VendorCar, "VendorCarRefer")
		for j,_ := range Car[i].VendorCar{
			db.Model(Car[i].VendorCar[j]).Related(&Car[i].VendorCar[j].Vendor, "VendorRefer")
			db.Model(Car[i].VendorCar[j].Vendor).Related(&Car[i].VendorCar[j].Vendor.VendorLocation, "VendorLocationRefer")
			for k,_ := range Car[i].VendorCar[j].Vendor.VendorLocation{
				db.Model(Car[i].VendorCar[j].Vendor.VendorLocation[k]).Related(&Car[i].VendorCar[j].Vendor.VendorLocation[k].Location, "LocationRefer")
			}
		}
	}
	return Car, err
}

func GetCarByLocation(location string)([]Car, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Car []Car
	db.Where("id IN (?)", db.Table("vendor_cars").Select("vendor_car_refer").Where("vendor_refer IN (?)", db.Table("vendors").Select("id").Where("id IN (?)", db.Table("vendor_locations").Select("vendor_location_refer").Where("location_refer IN (?)", db.Table("locations").Select("id").Where("city = ?", location).SubQuery()).SubQuery()).SubQuery()).SubQuery()).Find(&Car)
	for i,_ := range Car{
		db.Model(Car[i]).Related(&Car[i].VendorCar, "VendorCarRefer")
		for j,_ := range Car[i].VendorCar{
			db.Model(Car[i].VendorCar[j]).Related(&Car[i].VendorCar[j].Vendor, "VendorRefer")
			db.Model(Car[i].VendorCar[j].Vendor).Related(&Car[i].VendorCar[j].Vendor.VendorLocation, "VendorLocationRefer")
			for k,_ := range Car[i].VendorCar[j].Vendor.VendorLocation{
				db.Model(Car[i].VendorCar[j].Vendor.VendorLocation[k]).Related(&Car[i].VendorCar[j].Vendor.VendorLocation[k].Location, "LocationRefer")
			}
		}
	}
	return Car, err
}

func InsertCar(name string)(*Car, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	car := &Car{Name:name}

	db.Save(car)
	return car, err
}

func UpdateCar(id int, name string) (*Car, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var car Car
	db.Model(&car).Where("id = ?", id).Updates(map[string]interface{}{"name": name})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&car)
	return &car, nil
}

func RemoveCar(id int) (*Car, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	car := Car{Id: id}
	db.Where("id = ?", id).Find(&car)
	return &car, db.Delete(car).Error
}
