package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:relationship.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
