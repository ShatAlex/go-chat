package repository

import (
	"github.com/ShatAlex/chat"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user chat.User) (int, error)
	GetUserId(username, password_hash string) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
