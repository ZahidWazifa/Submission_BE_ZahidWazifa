package repository

import (
	"buy-sell-backend/internal/models"
	"database/sql"
)

// UserRepository adalah struktur untuk mengelola data pengguna
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository adalah fungsi untuk membuat instance UserRepository baru
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser adalah fungsi untuk menambahkan pengguna baru ke database
func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password)
	return err
}

// GetUser adalah fungsi untuk mengambil data pengguna berdasarkan ID
func (r *UserRepository) GetUser(id int) (*models.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser adalah fungsi untuk memperbarui data pengguna
func (r *UserRepository) UpdateUser(user *models.User) error {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.ID)
	return err
}

// DeleteUser adalah fungsi untuk menghapus pengguna dari database
func (r *UserRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}