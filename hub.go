package chat

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Rooms map[int]*Room
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[int]*Room),
	}
}

func (hub *Hub) PushNewUser(client *websocket.Conn, chatId int) []*websocket.Conn {
	room := hub.Rooms[chatId]
	room.Clients = append(room.Clients, client)
	fmt.Println(hub.Rooms[chatId])
	return room.Clients
}
