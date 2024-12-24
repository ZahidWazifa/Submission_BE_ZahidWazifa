// transaction_service.go

package service

import (
	"errors"
	"buy-sell-backend/internal/models"
	"buy-sell-backend/internal/repository"
)

// TransactionService adalah struktur yang menyimpan repositori transaksi
type TransactionService struct {
	repo repository.TransactionRepository
}

// NewTransactionService adalah fungsi untuk membuat instance baru dari TransactionService
func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

// CreateTransaction adalah metode untuk membuat transaksi baru
func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	if transaction == nil {
		return errors.New("transaksi tidak boleh kosong")
	}
	return s.repo.Create(transaction)
}

// GetTransaction adalah metode untuk mendapatkan transaksi berdasarkan ID
func (s *TransactionService) GetTransaction(id string) (*models.Transaction, error) {
	return s.repo.GetByID(id)
}

// GetAllTransactions adalah metode untuk mendapatkan semua transaksi
func (s *TransactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repo.GetAll()
}