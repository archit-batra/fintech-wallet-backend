package wallet

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Create wallet with zero balance
func (r *Repository) CreateWallet(userID int) error {

	query := `
		INSERT INTO wallets (user_id, balance)
		VALUES ($1, 0)
	`

	_, err := r.db.Exec(query, userID)
	return err
}

// Add money safely using transaction + row lock
func (r *Repository) AddBalance(userID int, amount int64) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var balance int64

	// Lock row
	err = tx.QueryRow(
		`SELECT balance FROM wallets WHERE user_id = $1 FOR UPDATE`,
		userID,
	).Scan(&balance)

	if err != nil {
		return err
	}

	newBalance := balance + amount

	_, err = tx.Exec(
		`UPDATE wallets SET balance = $1 WHERE user_id = $2`,
		newBalance,
		userID,
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *Repository) GetWallet(userID int) (Wallet, error) {

	var w Wallet

	err := r.db.QueryRow(
		`SELECT user_id, balance FROM wallets WHERE user_id = $1`,
		userID,
	).Scan(&w.UserID, &w.Balance)

	return w, err
}
