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

type Service struct {
	Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}
