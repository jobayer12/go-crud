package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := fmt.Sprintf("host=localhost user=postgres password=postgres dbname=golang port=5432 sslmode=disable TimeZone=Asia/Dhaka")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err = DB.AutoMigrate(&User{})
	if err != nil {
		return
	}
}
