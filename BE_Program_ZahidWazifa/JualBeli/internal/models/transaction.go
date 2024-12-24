package models

// Transaction represents a transaction in the buy and sell system
type Transaction struct {
	ID          int     `json:"id"`          // ID transaksi
	ProductID   int     `json:"product_id"`  // ID produk yang terlibat dalam transaksi
	UserID      int     `json:"user_id"`     // ID pengguna yang melakukan transaksi
	Quantity    int     `json:"quantity"`    // Jumlah barang yang ditransaksikan
	TotalPrice  float64 `json:"total_price"` // Total harga transaksi
	Status      string  `json:"status"`      // Status transaksi (e.g., completed, pending)
	CreatedAt   string  `json:"created_at"`  // Waktu transaksi dibuat
	UpdatedAt   string  `json:"updated_at"`  // Waktu transaksi diperbarui
}