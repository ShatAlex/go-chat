package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ShatAlex/chat"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func (h *WsHandler) joinRoom(c *gin.Context) {

	token := ""
	if cookie, err := c.Request.Cookie("AUTH"); err == nil {
		token = cookie.Value
	}

	var chats []chat.Chat

	userId, err := h.services.ParseToken(token)
	if err != nil {
		log.Print(err.Error())
		return
	}

	chats, err = h.services.GetUserChats(userId)
	if err != nil {
		log.Print(err.Error())
		return
	}

	chatId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid chatId")
		return
	}

	messages, err := h.services.GetMessages(chatId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var adminId int
	for _, v := range chats {
		if v.Id == chatId {
			adminId = v.Admin_id
		}
	}

	cl := &chat.Client{
		Conn:    nil,
		Message: make(chan *chat.Message, 10),
		Id:      userId,
		ChatId:  chatId,
	}

	m := &chat.Message{
		Chat_id: chatId,
		User_id: userId,
		Content: "User has joined the room",
	}

	c.HTML(http.StatusOK, "chat.html", gin.H{
		"chats":    chats,
		"messages": messages,
		"chatId":   chatId,
		"adminId":  adminId,
		"userId":   userId,
	})

	h.services.Websocket.RunRoomsMethods(cl, m)

}
