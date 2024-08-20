package database

import (
	"database/sql"

	"github.com/emiliosheinz/fc-ms-wallet-core/internal/entity"
)

type TransactioDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactioDB {
	return &TransactioDB{
		DB: db,
	}
}

func (t *TransactioDB) Create(transaction *entity.Transaction) error {
	stmt, error := t.DB.Prepare("INSERT INTO transactions (id, account_id_from, account_id_to, ammount, created_at) VALUES (?, ?, ?, ?, ?)")
	if error != nil {
		return error
	}
	defer stmt.Close()
	_, error = stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountTo.ID, transaction.Amount, transaction.CreatedAt)
	if error != nil {
		return error
	}
	return nil
}
