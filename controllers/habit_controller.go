package controllers

import (
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
	habits, err := ctrl.habitService.GetAllHabits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
