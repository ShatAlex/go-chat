package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ShatAlex/chat"
	"github.com/ShatAlex/chat/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WsHandler struct {
	services *service.Service
}

func NewWsHandler(ser *service.Service) *WsHandler {
	return &WsHandler{
		services: ser,
	}
}

func (h *WsHandler) joinRoom(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

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
		Conn:    conn,
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

// var clients []websocket.Conn

// func (h *Handler) WsEndpoint(w http.ResponseWriter, r *http.Request) {

// 	conn, _ := upgrader.Upgrade(w, r, nil)
// 	defer conn.Close()

// 	clients = append(clients, *conn)

// 	for {
// 		messageType, message, err := conn.ReadMessage()
// 		if err != nil || messageType == websocket.CloseMessage {
// 			return
// 		}

// 		for _, client := range clients {
// 			if err = client.WriteMessage(messageType, message); err != nil {
// 				log.Println(err)
// 				return
// 			}
// 		}
// 	}

// }
