package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var stringConnection = "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open(stringConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
