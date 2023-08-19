package service

import (
	"log"

	"github.com/ShatAlex/chat"
)

type WebsocketService struct {
	Hub *chat.Hub
}

func NewWebSocketService(hub *chat.Hub) *WebsocketService {
	return &WebSocketService{
		Hub: hub,
	}
}

func (s *WebsocketService) CreateRoom(chatId int, name string) error {
	s.Hub.Rooms[chatId] = &chat.Room{
		Id:     chatId,
		Name:   name,
		Clints: make(map[int]*chat.Client),
	}

	log.Print(s.Hub.Rooms[chatId])

	return nil
}
