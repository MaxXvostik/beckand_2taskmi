package database

import (
	"fmt"
	"log"
	"time"

	"exemple.com/beckand-2taskmi/model"
	"github.com/jinzhu/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=123456 dbname=beckand_2taskmi  sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.Tokens{}, &model.User{}, &model.Task{})
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("Database is unavailable , Wait for %d sec.\n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
