package repository

import (
	"database/sql"
	"errors"
	"buy-sell-backend/internal/models"
)

// ProductRepository adalah struktur untuk mengelola produk
type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository membuat instance baru dari ProductRepository
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Create menambahkan produk baru ke database
func (r *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, description, price) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price)
	if err != nil {
		return err
	}
	return nil
}

// GetByID mengambil produk berdasarkan ID
func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	product := &models.Product{}
	query := "SELECT id, name, description, price FROM products WHERE id = ?"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("produk tidak ditemukan")
		}
		return nil, err
	}
	return product, nil
}

// Update memperbarui informasi produk
func (r *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?"
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete menghapus produk dari database
func (r *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// GetAll mengambil semua produk dari database
func (r *ProductRepository) GetAll() ([]models.Product, error) {
	query := "SELECT id, name, description, price FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}