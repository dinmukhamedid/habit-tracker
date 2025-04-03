📌 Habit Tracker API
📝 Жоба сипаттамасы
Бұл жоба әдеттерді бақылау(habit-tracker) жүйесі үшін REST API ұсынады. Ол Golang, Gin, GORM және PostgreSQL көмегімен жасалған.

🛠 Қолданылған технологиялар
Golang — бағдарламалау тілі

Gin — веб-фреймворк

GORM — ORM (мәліметтер қорымен жұмыс істеу үшін)

PostgreSQL — мәліметтер қоры

🛠 Gin орнату
```
go get -u github.com/gin-gonic/gin
```
🌐 REST API
![image](https://github.com/user-attachments/assets/0daad456-f515-47ef-b864-e7c040f9d945)

🎛 Контроллер (controllers/user_controller.go)
Контроллер API-ға сұраныстарды қабылдап, оларды сервиске жібереді.

📌 Барлық пайдаланушыларды алу
```go
func (ctrl *UserController) GetAllUsers(c *gin.Context) {
    users, err := ctrl.userService.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}
```
Бұл GET /users сұранысын өңдейді және барлық пайдаланушыларды қайтарады.

📌 Email бойынша пайдаланушыны іздеу
```go
func (ctrl *UserController) GetUsersByEmail(c *gin.Context) {
    email := c.Query("email")
    users, err := ctrl.userService.FindUsersByEmail(email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}
```
Бұл GET /users/search?email=example@mail.com сұранысын өңдейді және email арқылы пайдаланушыны табады.

⚙ Сервис (services/user_service.go)
Сервис логикасы деректер базасынан мәліметтерді алады.

📌 Барлық пайдаланушыларды алу
```go
func (s *UserServiceImpl) GetAllUsers() ([]models.User, error) {
    return s.repo.GetAllUsers()
}
```
📌 Email бойынша пайдаланушыларды іздеу
```go
func (s *UserServiceImpl) FindUsersByEmail(email string) ([]models.User, error) {
    return s.repo.FindUsersByEmail(email)
}
```
🗂 Репозиторий (repository/user_repository.go)
Репозиторий GORM арқылы деректер базасымен жұмыс істейді.

📌 Email арқылы пайдаланушыны іздеу
```go
func (r *UserRepo) FindUsersByEmail(email string) ([]models.User, error) {
    var users []models.User
    result := config.DB.Where("email = ?", email).Find(&users)
    return users, result.Error
}
```
🛤 Маршруттар (routes/routes.go)
Бұл Gin көмегімен API маршруттарын орнатады.

```go
package routes

import (
	"habit-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserById)
	r.GET("/users/search", userController.GetUsersByEmail)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	return r
}
```






🔧 GORM қалай қостым?
1️⃣ GORM пакетін орнату
Алдымен gorm және PostgreSQL драйверін орнату керек.
Төмендегі команданы терминалда орындаңыз:
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```
2️⃣ Деректер базасына қосылу (config/database.go)
Бұл файлда GORM арқылы PostgreSQL-ға қосыламыз.

```go
package config

import (
    "fmt"
    "log"
    "habit-tracker/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=yourpassword dbname=habit_db port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Қате: мәліметтер базасына қосыла алмады", err)
    }

    DB = db
    Migrate()
    fmt.Println("✅ Мәліметтер базасына қосылу сәтті аяқталды!")
}

func Migrate() {
    err := DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Миграция қатесі:", err)
    }
    fmt.Println("✅ Миграция сәтті аяқталды!")
}
```
📌 Бұл код не істейді?

PostgreSQL-ға gorm.Open арқылы қосылады.

Қате болса, log.Fatal() арқылы шығарылады.

Migrate() функциясы User моделін базаға қосады.

3️⃣ User моделін құру (models/user.go)
GORM User моделін деректер базасымен байланысу үшін қолданамыз.

```go
package models

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"unique"`
    Email string `gorm:"unique"`
    Age   int
}
```
📌 Бұл код не істейді?

GORM аннотацияларын (gorm:"primaryKey" және gorm:"unique") қолданады.

ID — негізгі кілт.

Name және Email — бірегей мәндер.

4️⃣ Репозиторийдегі GORM сұраныстары (repository/user_repository.go)
Барлық пайдаланушыларды алу:
```go
func (r *UserRepo) GetAllUsers() ([]models.User, error) {
    var users []models.User
    result := config.DB.Find(&users)
    return users, result.Error
}
```
Бұл SELECT * FROM users; сұранысын орындайды.

ID арқылы пайдаланушыны табу:
```go
func (r *UserRepo) GetUserById(id uint) (models.User, error) {
    var user models.User
    result := config.DB.First(&user, id)
    return user, result.Error
}
```
Бұл SELECT * FROM users WHERE id = ? LIMIT 1; сұранысын орындайды.

Жаңа пайдаланушы қосу:
```go
func (r *UserRepo) CreateUser(user models.User) (models.User, error) {
    result := config.DB.Create(&user)
    return user, result.Error
}
```
Бұл INSERT INTO users (name, email, age) VALUES (?, ?, ?); сұранысын орындайды.

Email бойынша пайдаланушыны табу:
```go
func (r *UserRepo) FindUsersByEmail(email string) ([]models.User, error) {
    var users []models.User
    result := config.DB.Where("email = ?", email).Find(&users)
    return users, result.Error
}
```
Бұл SELECT * FROM users WHERE email = ?; сұранысын орындайды.
