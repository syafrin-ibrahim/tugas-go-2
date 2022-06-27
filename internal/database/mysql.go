package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/course?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cant connect to database")
	}

	return db

}
