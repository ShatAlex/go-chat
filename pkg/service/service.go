package service

import (
	"github.com/ShatAlex/chat"
	"github.com/ShatAlex/chat/pkg/repository"
)

type Authorization interface {
	CreateUser(user chat.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Chat interface {
	Create(name string, userId int) error
	GetUserChats(userId int) ([]chat.Chat, error)
	GetMessages(chatId, userId int) ([]chat.Message, error)
}

type Service struct {
	Authorization
	Chat
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
		Chat:          NewChatService(rep.Chat),
	}
}
