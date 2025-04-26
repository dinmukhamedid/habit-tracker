package controllers

import (
	"errors"
	"gorm.io/gorm"
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

func (ctrl *HabitController) GetAllHabits(c *gin.Context) {
	userID, _ := c.Get("user_id") // Middleware-дан алынады
	role, _ := c.Get("role")

	var habits []models.Habit
	if role == "admin" {
		// Admin барлық привычкаларды алады
		config.DB.Find(&habits)
	} else {
		// User тек өзінің привычкаларын алады
		config.DB.Where("user_id = ?", userID).Find(&habits)
	}

	c.JSON(http.StatusOK, habits)
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

	// ID-ді uint-ке түрлендіру
	id, err := strconv.ParseUint(habitID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}

	var habit models.Habit
	result := config.DB.First(&habit, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch habit"})
		return
	}

	// Рөлді тексеру
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	if role != "admin" && habit.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	// Привычканы өшіру
	config.DB.Delete(&habit)
	c.JSON(http.StatusOK, gin.H{"message": "Habit deleted successfully"})
}
