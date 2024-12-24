package config

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq" // Import PostgreSQL driver
)

// Konfigurasi database
const (
    host     = "localhost"
    port     = 5432
    user     = "yourusername"
    password = "yourpassword"
    dbname   = "buy_sell_db"
)

// Fungsi untuk menghubungkan ke database
func Connect() *sql.DB {
    // Membuat string koneksi
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // Menghubungkan ke database
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Gagal menghubungkan ke database: %v", err)
    }

    // Memeriksa koneksi
    err = db.Ping()
    if err != nil {
        log.Fatalf("Gagal melakukan ping ke database: %v", err)
    }

    log.Println("Berhasil terhubung ke database!")
    return db
}