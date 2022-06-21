package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"oneclick/server/model"
)

func initDB() *gorm.DB {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&",
		"root",
		"root",
		"127.0.0.1",
		"3306",
		"ss",
		"utf8",
	)
	DB, err := gorm.Open("mysql", args)
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&model.User{})
	return DB
	//db.SetMaxOpenConns(4)
	//db.SetMaxIdleConns(2)
}

func GetDB() *gorm.DB {
	db := initDB()
	return db
}
