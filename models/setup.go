package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Can't load configuration!")
	}
	DBName := os.Getenv("DB_NAME")
	DBUsername := os.Getenv("DB_USERNAME")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBPort := os.Getenv("DB_PORT")
	DBHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&mb4&parseTime=True&loc=Local", DBUsername, DBPassword, DBHost, DBPort, DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connection DB Failed")
		return
	}

	log.Println("Migrating Success")
	db.AutoMigrate(
		&Book{},
		&Employee{},
		&Visitor{},
		&User{})
	DB = db

}
