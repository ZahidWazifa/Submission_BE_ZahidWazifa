package repository

import (
    "database/sql"
    "errors"
    "github.com/yourusername/buy-sell-backend/internal/models"
)

// TransactionRepository adalah struktur untuk mengelola transaksi
type TransactionRepository struct {
    db *sql.DB
}

// NewTransactionRepository membuat instance baru dari TransactionRepository
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
    return &TransactionRepository{db: db}
}

// CreateTransaction menyimpan transaksi baru ke dalam database
func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
    query := "INSERT INTO transactions (user_id, product_id, amount) VALUES (?, ?, ?)"
    _, err := r.db.Exec(query, transaction.UserID, transaction.ProductID, transaction.Amount)
    if err != nil {
        return err
    }
    return nil
}

// GetTransactionByID mengambil transaksi berdasarkan ID
func (r *TransactionRepository) GetTransactionByID(id int) (*models.Transaction, error) {
    query := "SELECT id, user_id, product_id, amount FROM transactions WHERE id = ?"
    row := r.db.QueryRow(query, id)

    var transaction models.Transaction
    err := row.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Amount)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("transaksi tidak ditemukan")
        }
        return nil, err
    }
    return &transaction, nil
}

// GetAllTransactions mengambil semua transaksi dari database
func (r *TransactionRepository) GetAllTransactions() ([]models.Transaction, error) {
    query := "SELECT id, user_id, product_id, amount FROM transactions"
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var transactions []models.Transaction
    for rows.Next() {
        var transaction models.Transaction
        if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Amount); err != nil {
            return nil, err
        }
        transactions = append(transactions, transaction)
    }
    return transactions, nil
}