package db

import (
	"fmt"
	"log"
	"os"
	"todo_list/internal/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	log.Printf("Connecting to database with connection string: %s", connStr)

	var err error

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Successfully connected to database")

	DB.AutoMigrate(&models.Task{})
}
