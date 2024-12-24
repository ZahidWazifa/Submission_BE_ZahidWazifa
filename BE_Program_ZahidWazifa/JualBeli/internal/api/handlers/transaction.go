package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"buy-sell-backend/internal/models"
	"buy-sell-backend/internal/service"
)

// TransactionService untuk mengelola transaksi
var TransactionService = service.NewTransactionService()

// CreateTransaction menangani pembuatan transaksi baru
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}
	if err := TransactionService.CreateTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat transaksi"})
		return
	}
	c.JSON(http.StatusCreated, transaction)
}

// GetTransactions menangani pengambilan daftar transaksi
func GetTransactions(c *gin.Context) {
	transactions, err := TransactionService.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil transaksi"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}