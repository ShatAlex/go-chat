package service

import (
	"github.com/ShatAlex/chat"
	"github.com/ShatAlex/chat/pkg/repository"
)

type ChatService struct {
	repChat repository.Chat
}

func NewChatService(repChat repository.Chat) *ChatService {
	return &ChatService{repChat: repChat}
}

func (s *ChatService) Create(name string, userId int) error {
	return s.repChat.Create(name, userId)
}
func (s *ChatService) GetUserChats(userId int) ([]chat.Chat, error) {
	return s.repChat.GetUserChats(userId)
}

func (s *ChatService) GetMessages(chatId, userId int) ([]chat.Message, error) {
	return s.repChat.GetMessages(chatId, userId)
}
