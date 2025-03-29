package main

import (
	"habit-tracker/config"
	"habit-tracker/controllers"
	"habit-tracker/repository"
	"habit-tracker/routes"
	"habit-tracker/services"
)

func main() {
	config.ConnectDatabase()

	userRepo := &repository.UserRepo{}
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := routes.SetupRouter(userController)
	
	r.Run(":8080")
}
