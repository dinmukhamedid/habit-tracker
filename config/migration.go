package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	sqlBytes, err := ioutil.ReadFile("config/migration/1_init_tables.sql")
	if err != nil {
		log.Fatalf("Миграция файлын оқу қатесі: %v", err)
	}

	// SQL-ді бірнеше сұранысқа бөлу (егер бір файлда бірнеше сұраныс болса)
	queries := strings.Split(string(sqlBytes), ";")
	for _, query := range queries {
		q := strings.TrimSpace(query)
		if q != "" {
			if err := db.Exec(q).Error; err != nil {
				log.Fatalf("Миграция сұраныс қатесі: %v", err)
			}
		}
	}

	fmt.Println("🔧 SQL миграция сәтті орындалды.")
}
