package repository

import (
	"github.com/ShatAlex/chat"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user chat.User) (int, error)
	GetUserId(username, password_hash string) (int, error)
}

type Chat interface {
	Create(name string, userId int) error
	GetUserChats(userId int) ([]chat.Chat, error)
	GetMessages(chatId, userId int) ([]chat.Message, error)
}

type Repository struct {
	Authorization
	Chat
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Chat:          NewChatPostgres(db),
	}
}
