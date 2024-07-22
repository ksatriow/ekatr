package postgresql

import (
	"database/sql"

	"ekatr/internal/domain/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(u *user.User) error {
	query := "INSERT INTO users (username, password, email, type, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, u.Username, u.Password, u.Email, u.Type, u.CreatedAt)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	query := "SELECT id, username, password, email, type, created_at FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	var u user.User
	if err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Type, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindByUsername(username string) (*user.User, error) {
	query := "SELECT id, username, password, email, type, created_at FROM users WHERE username = $1"
	row := r.db.QueryRow(query, username)

	var u user.User
	if err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Type, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindByID(id int) (*user.User, error) {
	query := "SELECT id, username, password, email, type, created_at FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var u user.User
	if err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Type, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindAll() ([]*user.User, error) {
	query := "SELECT id, username, password, email, type, created_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Type, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
