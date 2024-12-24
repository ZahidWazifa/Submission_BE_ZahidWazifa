package service

import (
	"errors"
	"buy-sell-backend/internal/models"
	"buy-sell-backend/internal/repository"
)

// ProductService adalah struct yang menyediakan layanan untuk produk
type ProductService struct {
	repo repository.ProductRepository
}

// NewProductService adalah fungsi untuk membuat instance baru dari ProductService
func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// CreateProduct adalah metode untuk membuat produk baru
func (s *ProductService) CreateProduct(product *models.Product) error {
	if product.Name == "" {
		return errors.New("nama produk tidak boleh kosong")
	}
	return s.repo.Create(product)
}

// GetProduct adalah metode untuk mendapatkan produk berdasarkan ID
func (s *ProductService) GetProduct(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

// UpdateProduct adalah metode untuk memperbarui produk yang ada
func (s *ProductService) UpdateProduct(product *models.Product) error {
	if product.ID == 0 {
		return errors.New("ID produk tidak valid")
	}
	return s.repo.Update(product)
}

// DeleteProduct adalah metode untuk menghapus produk berdasarkan ID
func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}

// GetAllProducts adalah metode untuk mendapatkan semua produk
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}