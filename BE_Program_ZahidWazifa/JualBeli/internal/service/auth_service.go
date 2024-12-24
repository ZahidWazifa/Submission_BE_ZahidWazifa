// auth_service.go
package service

import (
	"errors"
	"buy-sell-backend/internal/models"
)

// AuthService adalah struct yang menyimpan dependensi untuk layanan otentikasi
type AuthService struct {
	userRepo UserRepository // repositori pengguna
}

// NewAuthService adalah konstruktor untuk AuthService
func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Login mengautentikasi pengguna berdasarkan email dan kata sandi
func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("pengguna tidak ditemukan")
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("kata sandi salah")
	}

	return user, nil
}

// Register mendaftarkan pengguna baru
func (s *AuthService) Register(user *models.User) error {
	existingUser, _ := s.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email sudah terdaftar")
	}

	return s.userRepo.Create(user)
}