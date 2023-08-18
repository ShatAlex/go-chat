package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients []websocket.Conn

func (h *Handler) WsEndpoint(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	clients = append(clients, *conn)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			return
		}

		for _, client := range clients {
			if err = client.WriteMessage(messageType, message); err != nil {
				log.Println(err)
				return
			}
		}
	}

}
