package models

// User represents the user model in the system
type User struct {
    ID       int    `json:"id"`       // ID pengguna
    Username string `json:"username"` // Nama pengguna
    Password string `json:"password"` // Kata sandi pengguna
    Email    string `json:"email"`    // Email pengguna
    CreatedAt string `json:"created_at"` // Tanggal pembuatan akun
    UpdatedAt string `json:"updated_at"` // Tanggal pembaruan akun
}