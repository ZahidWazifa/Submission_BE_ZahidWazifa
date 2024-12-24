package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"buy-sell-backend/internal/models"
	"buy-sell-backend/internal/service"
)

// UserHandler struct untuk menangani permintaan pengguna
type UserHandler struct {
	UserService service.UserService
}

// GetUser mengembalikan informasi pengguna berdasarkan ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Pengguna tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser memperbarui informasi pengguna
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data tidak valid"})
		return
	}
	updatedUser, err := h.UserService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui pengguna"})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}