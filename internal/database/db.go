package database

import (
	"log"

	"github.com/5afar/Junior_Backend_To-Do/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() {
	var err error
	connStr := "host=192.168.0.14 user=postgres dbname=todoapp sslmode=disable password=postgres"
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Создание таблиц
	DB.AutoMigrate(&models.Task{})
}
