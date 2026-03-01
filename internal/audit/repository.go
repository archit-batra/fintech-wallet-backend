package audit

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InsertLog(
	eventType string,
	fromUser int,
	toUser int,
	amount int64,
) error {

	query := `
		INSERT INTO audit_logs (event_type, from_user, to_user, amount)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(query, eventType, fromUser, toUser, amount)
	return err
}
