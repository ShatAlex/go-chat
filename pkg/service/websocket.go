package service

import (
	"github.com/ShatAlex/chat"
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
		Id:     chatId,
		Name:   name,
		Clints: make(map[int]*chat.Client),
	}

	return nil
}

func (s *WebSocketService) RunRoomsMethods(cl *chat.Client, m *chat.Message) {
	s.hub.Register <- cl
	s.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(s.hub)

}
