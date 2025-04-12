package main

import (
	"habit-tracker/config"
	"habit-tracker/controllers"
	"habit-tracker/repository"
	"habit-tracker/routes"
	"habit-tracker/services"
)

func main() {
	// Базаға қосылу
	config.ConnectDatabase()

	// User сервисі мен контроллерін инициализациялау
	userRepo := repository.NewUserRepo(config.DB)
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService()
	authController := controllers.NewAuthController(userService, authService)

	// Habit сервисі мен контроллерін инициализациялау
	habitRepo := repository.NewHabitRepo(config.DB)
	habitService := services.NewHabitService(habitRepo)
	habitController := controllers.NewHabitController(habitService)

	// Маршрутизаторды инициализациялау
	r := routes.SetupRouter(
		controllers.NewUserController(userService),
		habitController,
		authController,
	)

	// Серверді іске қосу
	r.Run(":8080")
}
