package models

import (
	"fmt"
	"github.com/NelsonMario/connection"
	"time"
)

type Blog struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	UserID    int        `gorm:"type:int;not null"`
	Title     string     `gorm:"type:varchar(100);not null"`
	Value     string     `gorm:"type:text;not null"`
	Image     string     `gorm:"type:text;not null"`
	Thumbnail string     `gorm:"type:varchar(100);not null"`
	Viewer int `gorm: "type:int;not null"`
}

func init(){
	db, err = connection.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Blog{})
}

func GetAllBlog() ([]Blog, error) {
	db, err = connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	var blogs []Blog

	db.Find(&blogs)

	return blogs, nil
}
func GetSingleBlog(id int) ([]Blog, error) {
	db, err = connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	var blog []Blog

	db.Where("id = ?", id).Find(&blog)

	return blog, nil
}

func InsertBlog(userId int, title string, value string, image string, thumbnail string, viewer int) (*Blog, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	blog := &Blog{
		UserID:	   userId,
		Title:     title,
		Value:     value,
		Image:     image,
		Thumbnail: thumbnail,
		Viewer: viewer,
	}
	db.Save(blog)

	return blog, nil
}

func UpdateBlog(id int, title string, thumbnail string, value string) (*Blog, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var blog Blog
	db.Model(&blog).Where("id = ?", id).Updates(map[string]interface{}{"title": title, "thumbnail": thumbnail, "value": value})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&blog)
	return &blog, nil
}

func RemoveBlog(id int) (*Blog, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	blog := Blog{ID: id}
	db.Where("id = ?", id).Delete(&blog)
	return &blog, db.Delete(blog).Error
}

func IncreaseViewer(id int, viewer int) (*Blog, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var blog Blog

	fmt.Print(viewer)
	db.Model(&blog).Where("id = ?", id).Updates(map[string]interface{}{"viewer": viewer})
	db.Where("id = ?", id).Find(&blog)
	return &blog, nil
}