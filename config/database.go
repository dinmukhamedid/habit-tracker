package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"habit-tracker/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=@zxcvbnm05 dbname=habit_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	DB = db

	Migrate()

	fmt.Println("Подключение к базе данных успешно!")
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
	fmt.Println("Миграция успешно выполнена!")
}
