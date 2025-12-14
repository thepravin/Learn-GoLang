package store

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "app_user:app_pass@tcp(127.0.0.1:3606)/demo?parseTime=true"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database...")
	}

	DB = database

	log.Println("Database connection established successfully.....")
}
