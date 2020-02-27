package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type Railroad struct {
	Id				int			`gorm:primary_key`
	Train 		Train		`gorm:"foreignkey:TrainRefer"`
	TrainRefer	int
	From			Station		`gorm:"foreignkey:FromRefer"`
	FromRefer		int
	To				Station		`gorm:"foreignkey:ToRefer"`
	ToRefer			int
	Departure		time.Time
	Arrival 		time.Time
	Duration		int
	Class           string
	Price			int				`gorm:"type:int"`
	Tax				int				`gorm:"type:int"`
	ServiceCharge	int				`gorm:"type:int"`
	RailroadRoutes	[]RailroadRoutes `gorm:"foreignkey:RailroadReferRoute"`
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
	db.AutoMigrate(&Railroad{}).AddForeignKey("train_refer", "trains(id)", "CASCADE", "CASCADE").AddForeignKey("from_refer", "stations(id)", "CASCADE", "CASCADE").AddForeignKey("to_refer", "stations(id)", "CASCADE", "CASCADE")
	GetRailroadBySchedule("Jakarta", "Bandung", "2020-02-01", "")
}

func InsertRailroad(trainRefer int, fromRefer int, toRefer int, depature time.Time, arrival time.Time, duration int, price int, tax int, serviceCharge int)(*Railroad, error){
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	Railroad := &Railroad{TrainRefer: trainRefer, FromRefer: fromRefer, ToRefer: toRefer, Departure: depature, Arrival: arrival, Duration: duration, Price: price, Tax: tax, ServiceCharge: serviceCharge}
	db.Save(Railroad)

	return Railroad, nil
}

func GetAllRailroad() ([]Railroad, error){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Railroad []Railroad

	db.Find(&Railroad)
	for i,_ := range Railroad{
		db.Model(Railroad[i]).Related(&Railroad[i].Train, "TrainRefer").Related(&Railroad[i].From, "FromRefer").Related(&Railroad[i].To, "ToRefer").Related(&Railroad[i].RailroadRoutes, "RailroadReferRoute")
		for j,_ := range Railroad[i].RailroadRoutes{
			db.Model(Railroad[i].RailroadRoutes[j]).Related(&Railroad[i].RailroadRoutes[j].Station, "StationRefer")
		}
	}
	return Railroad, err
}

func GetRailroadByFromAndTo(from string, to string) ([]Railroad, error){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Railroad []Railroad

	db.Where("from_refer IN (?) AND to_refer IN (?)", db.Table("stations").Select("id").Where("city = ?", from).SubQuery(), db.Table("stations").Select("id").Where("city = ?", to).SubQuery()).Find(&Railroad)
	for i,_ := range Railroad{
		db.Model(Railroad[i]).Related(&Railroad[i].Train, "TrainRefer").Related(&Railroad[i].From, "FromRefer").Related(&Railroad[i].To, "ToRefer").Related(&Railroad[i].RailroadRoutes, "RailroadReferRoute")
		for j,_ := range Railroad[i].RailroadRoutes{
			db.Model(Railroad[i].RailroadRoutes[j]).Related(&Railroad[i].RailroadRoutes[j].Station, "StationRefer")
		}
	}
	return Railroad, err
}


func GetRailroadByOneSchedule(fromCity string, toCity string, date string) ([]Railroad, error){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Railroad []Railroad

	db.Where("from_refer IN (?) AND to_refer IN (?) AND date_trunc('day', departure) = (?)", db.Table("stations").Select("id").Where("city = ?", fromCity).SubQuery(), db.Table("stations").Select("id").Where("city = ?", toCity).SubQuery(), date).Find(&Railroad)
	for i,_ := range Railroad{
		db.Model(Railroad[i]).Related(&Railroad[i].Train, "TrainRefer").Related(&Railroad[i].From, "FromRefer").Related(&Railroad[i].To, "ToRefer").Related(&Railroad[i].RailroadRoutes, "RailroadReferRoute")
		for j,_ := range Railroad[i].RailroadRoutes{
			db.Model(Railroad[i].RailroadRoutes[j]).Related(&Railroad[i].RailroadRoutes[j].Station, "StationRefer")
		}
	}
	fmt.Println(Railroad)
	return Railroad, err
}

func GetRailroadBySchedule(fromCity string, toCity string, fromDate string, toDate string) ([]Railroad, error){
	db, err = connection.ConnectDatabase()


	if toDate == ""{
		toDate = time.Now().Format("2006-01-02")
	}

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var Railroad []Railroad

	db.Where("((to_refer IN (?) AND from_refer IN (?)) OR (to_refer IN (?) AND from_refer IN (?))) AND ((date_trunc('day', departure) = (?)) OR (date_trunc('day', departure) = (?)))", db.Table("stations").Select("id").Where("city = ?", toCity).SubQuery(), db.Table("stations").Select("id").Where("city = ?", fromCity).SubQuery(), db.Table("stations").Select("id").Where("city = ?", fromCity).SubQuery(), db.Table("stations").Select("id").Where("city = ?", toCity).SubQuery(), fromDate ,toDate).Find(&Railroad)
	for i,_ := range Railroad{
		db.Model(Railroad[i]).Related(&Railroad[i].Train, "TrainRefer").Related(&Railroad[i].From, "FromRefer").Related(&Railroad[i].To, "ToRefer").Related(&Railroad[i].RailroadRoutes, "RailroadReferRoute")
		for j,_ := range Railroad[i].RailroadRoutes{
			db.Model(Railroad[i].RailroadRoutes[j]).Related(&Railroad[i].RailroadRoutes[j].Station, "StationRefer")
		}
	}
	fmt.Println(Railroad)
	return Railroad, err
}

func UpdateRailroad(id int, trainRefer int, fromRefer int, toRefer int, depature time.Time, arrival time.Time, duration int, price int, tax int, serviceCharge int) (*Railroad, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var railRoad Railroad
	db.Model(&railRoad).Where("id = ?", id).Updates(map[string]interface{}{"train_refer": trainRefer, "from_refer": fromRefer, "to_refer": toRefer, "departure": depature, "arrival": arrival, "duration": duration, "price": price, "tax": tax, "service_charge": serviceCharge})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&railRoad)
	return &railRoad, nil
}

func RemoveRailroad(id int) (*Railroad, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	railRoad := Railroad{Id: id}
	db.Where("id = ?", id).Find(&railRoad)
	return &railRoad, db.Delete(railRoad).Error
}


