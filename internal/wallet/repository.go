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

func (r *Repository) Transfer(fromID, toID int, amount int64) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	// Always lock rows in consistent order to avoid deadlocks
	first := fromID
	second := toID

	if fromID > toID {
		first = toID
		second = fromID
	}

	// Lock both wallet rows
	_, err = tx.Exec(
		`SELECT user_id FROM wallets WHERE user_id IN ($1,$2) ORDER BY user_id FOR UPDATE`,
		first,
		second,
	)
	if err != nil {
		return err
	}

	var fromBalance int64

	// Get sender balance
	err = tx.QueryRow(
		`SELECT balance FROM wallets WHERE user_id = $1`,
		fromID,
	).Scan(&fromBalance)
	if err != nil {
		return err
	}

	if fromBalance < amount {
		return sql.ErrNoRows
	}

	// Deduct sender
	_, err = tx.Exec(
		`UPDATE wallets SET balance = balance - $1 WHERE user_id = $2`,
		amount,
		fromID,
	)
	if err != nil {
		return err
	}

	// Add receiver
	_, err = tx.Exec(
		`UPDATE wallets SET balance = balance + $1 WHERE user_id = $2`,
		amount,
		toID,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}
