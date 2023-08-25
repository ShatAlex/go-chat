package chat

import (
	"time"

	"github.com/gorilla/websocket"
)

type Chat struct {
	Id       int    `json:"_"`
	Name     string `json:"name" binding:"required"`
	Admin_id int    `json:"admin_id" binding:"required"`
}

type Room struct {
	Id      int               `json:"_"`
	Name    string            `json:"name"`
	Clients []*websocket.Conn `json:"clients"`
}

type Message struct {
	Id       int       `json:"_"`
	Chat_id  int       `json:"chat_id" binding:"required"`
	User_id  int       `json:"user_id" binding:"required"`
	Datetime time.Time `json:"datetime" binding:"required"`
	Content  string    `json:"content" binding:"required"`
}
