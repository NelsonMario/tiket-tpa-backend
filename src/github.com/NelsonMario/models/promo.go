package models

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/middleware"
	"time"
)

type Promo struct{
	Id			int			`gorm:primary_key`
	Name string `gorm:"type:varchar(100)"`
	StartDate time.Time
	EndDate time.Time
	PromoDetail []PromoDetail `gorm:"foreignkey:PromoRefer"`
	Description string
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
	db.AutoMigrate(&Promo{})
}

func GetAllPromo() ([]Promo, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Promo []Promo

	db.Find(&Promo)
	for i,_ := range Promo{
		db.Model(Promo[i]).Related(&Promo[i].PromoDetail, "PromoRefer")
	}

	return Promo, err
}

func GetPromoById(id int) ([]Promo, error){
	db, err = connection.ConnectDatabase()
_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Promo []Promo

	db.Where("id = ?", id).Find(&Promo)
	for i,_ := range Promo{
		db.Model(Promo[i]).Related(&Promo[i].PromoDetail, "PromoRefer")
	}

	return Promo, err
}

