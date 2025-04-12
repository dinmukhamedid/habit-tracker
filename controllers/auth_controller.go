package controllers

import (
	"habit-tracker/models"
	"habit-tracker/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService services.UserService
	authService services.AuthService
}

func NewAuthController(userService services.UserService, authService services.AuthService) *AuthController {
	return &AuthController{userService, authService}
}
func (ctrl *AuthController) Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Default role: "user"
	input.Role = "user"

	user, err := ctrl.userService.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists or invalid data"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userService.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := ctrl.authService.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
