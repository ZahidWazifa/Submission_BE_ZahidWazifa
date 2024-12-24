package models

// Product adalah model untuk barang yang dijual
type Product struct {
	ID          int     `json:"id"`          // ID unik untuk produk
	Name        string  `json:"name"`        // Nama produk
	Description string  `json:"description"` // Deskripsi produk
	Price       float64 `json:"price"`       // Harga produk
	Quantity    int     `json:"quantity"`    // Jumlah produk yang tersedia
}

// NewProduct adalah fungsi untuk membuat produk baru
func NewProduct(id int, name string, description string, price float64, quantity int) *Product {
	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}