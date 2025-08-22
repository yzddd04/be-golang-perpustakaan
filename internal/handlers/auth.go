package handlers

import (
	"net/http"

	"library-management-system/internal/models"
	"library-management-system/internal/repository"
	"library-management-system/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userRepo *repository.UserRepository
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userRepo: repository.NewUserRepository(),
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func Register(c *gin.Context) {
	handler := NewAuthHandler()
	
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	existingUser, _ := handler.userRepo.GetByUsername(req.Username)
	if existingUser != nil {
		utils.ErrorResponse(c, http.StatusConflict, "Username already exists")
		return
	}

	existingEmail, _ := handler.userRepo.GetByEmail(req.Email)
	if existingEmail != nil {
		utils.ErrorResponse(c, http.StatusConflict, "Email already exists")
		return
	}

	if req.Role == "" {
		req.Role = "user"
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	if err := handler.userRepo.Create(user); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SuccessResponse(c, "User registered successfully", user)
}

func Login(c *gin.Context) {
	handler := NewAuthHandler()
	
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	user, err := handler.userRepo.GetByUsername(req.Username)
	if err != nil {
		utils.UnauthorizedResponse(c, "Invalid credentials")
		return
	}

	if !user.CheckPassword(req.Password) {
		utils.UnauthorizedResponse(c, "Invalid credentials")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := LoginResponse{
		Token: token,
		User:  *user,
	}

	utils.SuccessResponse(c, "Login successful", response)
}
