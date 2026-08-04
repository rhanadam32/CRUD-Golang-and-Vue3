package database

import (
	"fmt"
	"log"
	"santri-koding-trial/backend-golang/config"
	"santri-koding-trial/backend-golang/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbUser := config.GetEnv("DB_USER", "root")
	dbName := config.GetEnv("DB_NAME", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306")
	dbPass := config.GetEnv("DB_PASS", "root")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect Database:", err)
	}

	fmt.Println("Database connected successfully!")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migrated successfully!")
}
