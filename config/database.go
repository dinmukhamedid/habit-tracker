package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"habit-tracker/models"
)

var DB *gorm.DB
var JWTSecret string // JWT секретін сақтау үшін

func ConnectDatabase() {
	// .env файлын жүктеу
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке env. файла")
	}

	// DSN (Data Source Name) жасау
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// 🔌 GORM арқылы базаға қосылу
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	DB = db

	// JWT секретін жүктеу
	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatal("JWT_SECRET анықталмаған")
	}

	// Миграцияларды орындау
	Migrate()

	fmt.Println("Подключение к базе данных успешно!")
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Habit{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
	fmt.Println("Миграция успешно выполнена!")
}
