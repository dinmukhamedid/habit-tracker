package routes

import (
	"github.com/gin-gonic/gin"
	"habit-tracker/controllers"
	"habit-tracker/middleware"
)

func SetupRouter(
	userController *controllers.UserController,
	habitController *controllers.HabitController,
	authController *controllers.AuthController,
) *gin.Engine {
	r := gin.Default()

	// User маршруттары
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserById)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	// Habit маршруттары
	r.GET("/habits", habitController.GetAllHabits)
	r.GET("/habits/user/:userId", habitController.GetHabitsByUserId)
	r.POST("/habits", habitController.CreateHabit)

	// Auth маршруттары
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	// ✅ Middleware-ді тек /habits маршруттарына қолдану
	habits := r.Group("/habits")
	habits.Use(middleware.AuthMiddleware()) // AuthService-ті middleware-ге беру
	{
		habits.GET("/", habitController.GetAllHabits)
		habits.POST("/", habitController.CreateHabit)
	}

	return r
}
