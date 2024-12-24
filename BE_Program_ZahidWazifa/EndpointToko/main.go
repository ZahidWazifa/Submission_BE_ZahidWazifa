package main

import (
	"task3/config"
	"task3/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi koneksi database
	db := config.InitDB()
	
	// Inisialisasi router Gin
	r := gin.Default()
	
	// Setup routes
	routes.SetupRoutes(r, db)
	
	// Jalankan server pada port 8080
	r.Run(":8080")
}
