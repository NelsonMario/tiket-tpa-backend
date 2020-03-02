package models

import (
	"github.com/NelsonMario/connection"
	"time"
)

type PromoDetail struct{
	Id			int			`gorm:primary_key`
	PromoRefer int
	Disc int
	PromoCode string `gorm:"type:varchar(100)"`
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
	db.AutoMigrate(&PromoDetail{}).AddForeignKey("promo_refer", "promos(id)", "CASCADE", "CASCADE")
}


