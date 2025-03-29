package routes

import (
	"habit-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
	r := gin.Default()
	
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserById)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	return r
}
