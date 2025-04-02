package services

import (
	"habit-tracker/models"
	"habit-tracker/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	FindUsersByEmail(email string) ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserServiceImpl) GetUserById(id uint) (models.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UserServiceImpl) FindUsersByEmail(email string) ([]models.User, error) {
	return s.repo.FindUsersByEmail(email)
}

func (s *UserServiceImpl) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserServiceImpl) UpdateUser(user models.User) (models.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *UserServiceImpl) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
