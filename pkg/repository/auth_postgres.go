package repository

import (
	"fmt"

	"github.com/ShatAlex/chat"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user chat.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserId(username, password_hash string) (int, error) {
	var userId int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", usersTable)
	err := r.db.Get(&userId, query, username, password_hash)
	return userId, err
}
