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
	Create(name string, userId int) (int, error)
	GetUserChats(userId int) ([]chat.Chat, error)
	GetMessages(chatId, userId int) ([]chat.Message, error)
	GetUserIdByUsername(username string) (int, error)
	AddUser(chatId, userId int) error
	CreateMessage(userId, chatId int, content string) error
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
