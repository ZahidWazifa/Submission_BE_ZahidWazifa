package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"buy-sell-backend/internal/models"
	"buy-sell-backend/internal/service"
)

// ProductService adalah service untuk mengelola produk
var ProductService = service.NewProductService()

// CreateProduct menangani permintaan untuk membuat produk baru
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}
	if err := ProductService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat produk"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// GetProducts menangani permintaan untuk mendapatkan daftar produk
func GetProducts(c *gin.Context) {
	products, err := ProductService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan produk"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// UpdateProduct menangani permintaan untuk memperbarui produk
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}
	if err := ProductService.UpdateProduct(id, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui produk"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// DeleteProduct menangani permintaan untuk menghapus produk
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := ProductService.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus produk"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}