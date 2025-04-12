package repository

import (
	"gorm.io/gorm"
	"habit-tracker/config"
	"habit-tracker/models"
)

type HabitRepository interface {
	GetAllHabits() ([]models.Habit, error)
	CreateHabit(habit models.Habit) (models.Habit, error)
	GetHabitsByUserId(userId uint) ([]models.Habit, error)
}

type HabitRepo struct {
	DB *gorm.DB
}

func NewHabitRepo(db *gorm.DB) *HabitRepo {
	return &HabitRepo{DB: db}
}

func (r *HabitRepo) GetAllHabits() ([]models.Habit, error) {
	var habits []models.Habit
	result := config.DB.Find(&habits)
	return habits, result.Error
}

func (r *HabitRepo) CreateHabit(habit models.Habit) (models.Habit, error) {
	result := config.DB.Create(&habit)
	return habit, result.Error
}

func (r *HabitRepo) GetHabitsByUserId(userId uint) ([]models.Habit, error) {
	var habits []models.Habit
	result := config.DB.Where("user_id = ?", userId).Find(&habits)
	return habits, result.Error
}
