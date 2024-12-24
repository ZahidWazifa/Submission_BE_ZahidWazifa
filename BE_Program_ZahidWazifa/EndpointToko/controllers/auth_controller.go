package controllers

import (
	"task3/models"
	"task3/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register menangani pendaftaran user baru
func (ac *AuthController) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	if err := ac.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Username atau email sudah digunakan"})
		return
	}

	c.JSON(201, gin.H{"message": "User berhasil didaftarkan"})
}

// Login menangani proses login
func (ac *AuthController) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := ac.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Username atau password salah"})
		return
	}

	if err := user.ValidatePassword(input.Password); err != nil {
		c.JSON(401, gin.H{"error": "Username atau password salah"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal membuat token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
