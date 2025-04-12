package services

import (
	"habit-tracker/config"
	"habit-tracker/models"
	"habit-tracker/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
	Register(user models.User) (models.User, error)    // Қосыңыз
	Login(email, password string) (models.User, error) // Қосыңыз
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

func (s *UserServiceImpl) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserServiceImpl) UpdateUser(user models.User) (models.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *UserServiceImpl) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

// Парольді хэштеу және қолданушыны сақтау
func (s *UserServiceImpl) Register(user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)
	err = config.DB.Create(&user).Error
	return user, err
}

// Email арқылы қолданушыны табу және парольді тексеру
func (s *UserServiceImpl) Login(email, password string) (models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err
}
