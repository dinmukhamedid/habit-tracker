package repository

import (
	"habit-tracker/config"
	"habit-tracker/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

type UserRepo struct{}

func (r *UserRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func (r *UserRepo) GetUserById(id uint) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return user, result.Error
}

func (r *UserRepo) CreateUser(user models.User) (models.User, error) {
	result := config.DB.Create(&user)
	return user, result.Error
}

func (r *UserRepo) UpdateUser(user models.User) (models.User, error) {
	result := config.DB.Save(&user)
	return user, result.Error
}

func (r *UserRepo) DeleteUser(id uint) error {
	result := config.DB.Delete(&models.User{}, id)
	return result.Error
}
