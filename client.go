package chat

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	Messagee chan *Message
	Id       int `json:"Id"`
	ChatId   int `json:"chatId"`
}
