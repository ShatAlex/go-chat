package main

import (
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

// func (server *Server) wsEndpoint(w http.ResponseWriter, r *http.Request) {

// 	conn, _ := upgrader.Upgrade(w, r, nil)
// 	defer conn.Close()

// 	server.clients[conn] = true
// 	defer delete(server.clients, conn)

// 	for {
// 		messageType, message, err := conn.ReadMessage()
// 		if err != nil || messageType == websocket.CloseMessage {
// 			break
// 		}

// 		if err := conn.WriteMessage(messageType, message); err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		go server.handleMessage(message)
// 	}

// }

// func (server *Server) WriteMessage(message []byte) {
// 	for conn := range server.clients {
// 		conn.WriteMessage(websocket.TextMessage, message)
// 	}
// }
