package config

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"task3/models"
)

// InitDB melakukan inisialisasi koneksi database
func InitDB() *gorm.DB {
	// Konfigurasi database
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Membuat DSN string untuk koneksi MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName)

	// Membuat koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi ke database: " + err.Error())
	}

	// Auto Migrate models
	db.AutoMigrate(&models.User{}, &models.Product{})

	return db
}
