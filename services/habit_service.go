package services

import (
	"fmt"
	"habit-tracker/config"
	"habit-tracker/models"
	"habit-tracker/repository"
)

type HabitService interface {
	GetAllHabits() ([]models.Habit, error)
	CreateHabit(habit models.Habit) (models.Habit, error)
	GetHabitsByUserId(userId uint) ([]models.Habit, error)
}
type HabitServiceImpl struct {
	repo repository.HabitRepository
}

func NewHabitService(repo repository.HabitRepository) *HabitServiceImpl {
	return &HabitServiceImpl{repo: repo}
}

func (s *HabitServiceImpl) GetAllHabits() ([]models.Habit, error) {
	return s.repo.GetAllHabits()
}
func (s *HabitServiceImpl) CreateHabit(habit models.Habit) (models.Habit, error) {
	// Қолданушының бар-жоғын тексеру
	var user models.User
	if err := config.DB.First(&user, habit.UserID).Error; err != nil {
		return models.Habit{}, fmt.Errorf("user not found")
	}

	// Егер бар болса, әдетті қосамыз
	result := config.DB.Create(&habit)
	return habit, result.Error
}

func (s *HabitServiceImpl) GetHabitsByUserId(userId uint) ([]models.Habit, error) {
	return s.repo.GetHabitsByUserId(userId)
}
