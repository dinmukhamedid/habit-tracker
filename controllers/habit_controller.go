package controllers

import (
	"habit-tracker/config"
	"habit-tracker/models"
	"habit-tracker/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HabitController struct {
	habitService services.HabitService
}

func NewHabitController(service services.HabitService) *HabitController {
	return &HabitController{habitService: service}
}

func (ctrl *HabitController) CreateHabit(c *gin.Context) {
	var habit models.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdHabit, err := ctrl.habitService.CreateHabit(habit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdHabit)
}

func (h *HabitController) GetAllHabits(c *gin.Context) {
	// Пайдаланушының ID-ін және рөлін алу
	userID, _ := c.Get("userID") // Middleware-ден алынады
	role, _ := c.Get("role")

	var habits []models.Habit
	if role == "admin" {
		// Admin барлық привычкаларды алады
		config.DB.Find(&habits)
	} else {
		// User тек өзінің привычкаларын алады
		config.DB.Where("user_id = ?", userID).Find(&habits)
	}

	c.JSON(200, habits)
}

func (ctrl *HabitController) GetHabitsByUserId(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Жарамсыз user ID"})
		return
	}

	habits, err := ctrl.habitService.GetHabitsByUserId(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habits)
}

func (h *HabitController) DeleteHabit(c *gin.Context) {
	habitID := c.Param("id")
	userID, _ := c.Get("userID") // Middleware-ден алынады
	role, _ := c.Get("role")

	var habit models.Habit
	if err := config.DB.First(&habit, habitID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Habit not found"})
		return
	}

	// Рөлді тексеру
	if role != "admin" && habit.UserID != userID.(uint) {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	// Привычканы өшіру
	config.DB.Delete(&habit)
	c.JSON(200, gin.H{"message": "Habit deleted successfully"})
}
