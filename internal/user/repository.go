package user

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(name, email string) (User, error) {

	var u User

	query := `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id, name, email
	`

	err := r.db.QueryRow(query, name, email).
		Scan(&u.ID, &u.Name, &u.Email)

	return u, err
}

func (r *Repository) GetUser(id string) (User, error) {

	var u User

	query := `
		SELECT id, name, email
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(query, id).
		Scan(&u.ID, &u.Name, &u.Email)

	return u, err
}
