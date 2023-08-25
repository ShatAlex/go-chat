package service

import (
	"fmt"

	"github.com/ShatAlex/chat"
	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	hub *chat.Hub
}

func NewWebSocketService(hub *chat.Hub) *WebSocketService {
	return &WebSocketService{
		hub: hub,
	}
}

func (s *WebSocketService) CreateRoom(chatId int, name string) error {
	s.hub.Rooms[chatId] = &chat.Room{
		Id:      chatId,
		Name:    name,
		Clients: []*websocket.Conn{},
	}
	fmt.Println(s.hub.Rooms)

	return nil
}

func (s *WebSocketService) PushNewUser(client *websocket.Conn, chatId int) []*websocket.Conn {
	return s.hub.PushNewUser(client, chatId)
}
