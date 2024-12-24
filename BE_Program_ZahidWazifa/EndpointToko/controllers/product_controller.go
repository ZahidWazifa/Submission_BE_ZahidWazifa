package controllers

import (
	"task3/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

type ProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
}

// Create menambah produk baru
func (pc *ProductController) Create(c *gin.Context) {
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		UserID:      userID.(uint),
	}

	if err := pc.DB.Create(&product).Error; err != nil {
		c.JSON(400, gin.H{"error": "Gagal menambah produk"})
		return
	}

	c.JSON(201, gin.H{"data": product})
}

// GetAll mengambil semua produk
func (pc *ProductController) GetAll(c *gin.Context) {
	var products []models.Product
	if err := pc.DB.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal mengambil data produk"})
		return
	}

	c.JSON(200, gin.H{"data": products})
}

// GetByID mengambil produk berdasarkan ID
func (pc *ProductController) GetByID(c *gin.Context) {
	var product models.Product
	if err := pc.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	c.JSON(200, gin.H{"data": product})
}

// Update memperbarui data produk
func (pc *ProductController) Update(c *gin.Context) {
	var product models.Product
	if err := pc.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	userID := c.GetUint("user_id")
	if product.UserID != userID {
		c.JSON(403, gin.H{"error": "Tidak memiliki akses"})
		return
	}

	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updateData := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}

	if err := pc.DB.Model(&product).Updates(updateData).Error; err != nil {
		c.JSON(400, gin.H{"error": "Gagal memperbarui produk"})
		return
	}

	c.JSON(200, gin.H{"data": product})
}

// Delete menghapus produk
func (pc *ProductController) Delete(c *gin.Context) {
	var product models.Product
	if err := pc.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	userID := c.GetUint("user_id")
	if product.UserID != userID {
		c.JSON(403, gin.H{"error": "Tidak memiliki akses"})
		return
	}

	if err := pc.DB.Delete(&product).Error; err != nil {
		c.JSON(400, gin.H{"error": "Gagal menghapus produk"})
		return
	}

	c.JSON(200, gin.H{"message": "Produk berhasil dihapus"})
}
