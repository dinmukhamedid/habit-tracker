# habit-tracker
 
📌 Жоба туралы
Бұл жоба — Go тілінде жазылған REST API, ол Gin веб-фреймворкін және GORM ORM-ін пайдаланып, PostgreSQL дерекқорымен жұмыс істейді.

Бұл API не істей алады?
✅ Пайдаланушыларды тіркеу
✅ Барлық пайдаланушыларды алу
✅ Email бойынша пайдаланушыны іздеу
✅ Пайдаланушыларды жасына қарай сұрыптау
✅ Пайдаланушыларды жаңарту және жою

🛠 Қолданылған технологиялар
Go — бағдарламалау тілі

Gin — HTTP сервер жасауға арналған веб-фреймворк

GORM — PostgreSQL-пен жұмыс істеуге арналған ORM

PostgreSQL — дерекқор

📌 REST API (Эндпоинттер)
HTTP әдісі	                    URL	Сипаттама	                                 Деректер (JSON)
GET	/users	                    Барлық пайдаланушыларды алу	                         ❌
GET	/users/:id	                Пайдаланушыны ID бойынша алу	                       ❌
GET	/users/search?email={email}	Email бойынша іздеу (жасына қарай сұрыптау)	         ❌
POST	/users	                  Жаңа пайдаланушы қосу	                               ✅ { "name": "Dinmukhamed", "email": "dinmukhamed@example.com", "age": 19 }
PUT	/users/:id	                Пайдаланушы деректерін жаңарту	                     ✅ { "name": "Arman", "email": "Arman@example.com", "age": 20 }
DELETE	/users/:id	            Пайдаланушыны жою	                                   ❌



💾 GORM (ORM) қалай жұмыс істейді?
GORM — бұл Go тіліндегі ORM (Object-Relational Mapping), ол SQL сұраныстарын автоматты түрде жасауға көмектеседі.

📌 Дерекқорға қосылу

Файл: config/database.go

```go
package config

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=habit_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Қате: дерекқорға қосылу мүмкін болмады!", err)
	}

	DB = db
	Migrate()
	fmt.Println("✅ Дерекқорға сәтті қосылдық!")
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{}) // users кестесін құру
	if err != nil {
		log.Fatal("Қате: миграция орындалмады!", err)
	}
	fmt.Println("✅ Миграция сәтті аяқталды!")
}

```
📌 Пайдаланушы моделі
Файл: models/user.go
```go
package models

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `json:"name"`
    Email string `json:"email" gorm:"unique"`
    Age   int    `json:"age"`
}

