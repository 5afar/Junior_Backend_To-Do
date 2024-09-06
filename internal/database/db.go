package database

import (
	"log"
	"fmt"

	"github.com/5afar/Junior_Backend_To-Do/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init(host, user, dbname, password string) {
	var err error
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",host,user,dbname,password)
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Создание таблиц
	DB.AutoMigrate(&models.Task{})
}
