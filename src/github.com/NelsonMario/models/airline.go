package models

type Airplane struct{

	ID int `gorm: "primary_key"`
	Name string `gorm: "type: varchar(100);not null"`
	Code string `gorm: "type: varchar(100);not null"`
	Seat int `gorm: "type int; not null"`
	Price int `gorm: "type int; not null"`
	

}

