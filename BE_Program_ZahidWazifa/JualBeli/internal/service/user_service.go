// user_service.go
package service

import (
	"errors"
	"buy-sell-backend/internal/models"
	"buy-sell-backend/internal/repository"
)

// UserService adalah struktur yang menyimpan repositori pengguna
type UserService struct {
	repo repository.UserRepository
}

// NewUserService adalah konstruktor untuk UserService
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser mengembalikan informasi pengguna berdasarkan ID
func (s *UserService) GetUser(id string) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser memperbarui informasi pengguna
func (s *UserService) UpdateUser(user *models.User) error {
	if user.ID == "" {
		return errors.New("ID pengguna tidak boleh kosong")
	}
	return s.repo.Update(user)
}

// DeleteUser menghapus pengguna berdasarkan ID
func (s *UserService) DeleteUser(id string) error {
	if id == "" {
		return errors.New("ID pengguna tidak boleh kosong")
	}
	return s.repo.Delete(id)
}