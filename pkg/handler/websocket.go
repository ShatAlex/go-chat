package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *WsHandler) wshandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	chatId, err := strconv.Atoi(strings.Split(r.URL.RawQuery, "=")[1])
	if err != nil {
		log.Println(err.Error())
		return
	}

	userId, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	if err != nil {
		log.Println(err.Error())
		return
	}

	clients := h.services.Websocket.PushNewUser(conn, chatId)

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("%s - user_id: %d on chat_id: %d send: %s\n", conn.RemoteAddr(), userId, chatId, string(msg))

		for _, client := range clients {
			if err = client.WriteMessage(msgType, msg); err != nil {
				return
			}
		}

	}
}
