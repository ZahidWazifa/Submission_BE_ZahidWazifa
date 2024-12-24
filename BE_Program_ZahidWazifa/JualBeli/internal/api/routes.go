package api

import (
    "github.com/gin-gonic/gin"
    "buy-sell-backend/internal/api/handlers"
)

// SetupRoutes mengatur rute untuk aplikasi
func SetupRoutes(router *gin.Engine) {
    // Rute untuk otentikasi
    authGroup := router.Group("/auth")
    {
        authGroup.POST("/login", handlers.LoginHandler)   // Rute untuk login
        authGroup.POST("/register", handlers.RegisterHandler) // Rute untuk registrasi
    }

    // Rute untuk produk
    productGroup := router.Group("/products")
    {
        productGroup.GET("/", handlers.GetProductsHandler)         // Rute untuk mendapatkan semua produk
        productGroup.POST("/", handlers.CreateProductHandler)      // Rute untuk membuat produk baru
        productGroup.PUT("/:id", handlers.UpdateProductHandler)    // Rute untuk memperbarui produk
        productGroup.DELETE("/:id", handlers.DeleteProductHandler)  // Rute untuk menghapus produk
    }

    // Rute untuk transaksi
    transactionGroup := router.Group("/transactions")
    {
        transactionGroup.GET("/", handlers.GetTransactionsHandler) // Rute untuk mendapatkan semua transaksi
        transactionGroup.POST("/", handlers.CreateTransactionHandler) // Rute untuk membuat transaksi baru
    }

    // Rute untuk pengguna
    userGroup := router.Group("/users")
    {
        userGroup.GET("/:id", handlers.GetUserHandler)         // Rute untuk mendapatkan informasi pengguna
        userGroup.PUT("/:id", handlers.UpdateUserHandler)      // Rute untuk memperbarui profil pengguna
    }
}