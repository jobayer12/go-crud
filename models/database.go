package models

import (
	"fmt"
	"github.com/jobayer12/go-crud/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	config, err := environment.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load environment variables ", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

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
