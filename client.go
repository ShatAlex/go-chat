package chat

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Message chan *Message
	Id      int `json:"Id"`
	ChatId  int `json:"chatId"`
}

func (cl *Client) WriteMessage() {
	defer cl.Conn.Close()

	for {
		message, ok := <-cl.Message
		if !ok {
			return
		}

		cl.Conn.WriteJSON(message)
	}
}

func (cl *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- cl
		cl.Conn.Close()
	}()

	for {
		_, message, err := cl.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %s", err)
			}

			break
		}

		msg := &Message{
			User_id: cl.Id,
			Chat_id: cl.ChatId,
			Content: string(message),
		}

		hub.Broadcast <- msg
	}
}
