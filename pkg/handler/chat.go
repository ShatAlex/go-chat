package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ShatAlex/chat"
	"github.com/gin-gonic/gin"
)

func (h *Handler) homePage(c *gin.Context) {

	token := ""
	if cookie, err := c.Request.Cookie("AUTH"); err == nil {
		token = cookie.Value
	}

	var chats []chat.Chat

	if token != "" {
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

	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"token": token,
		"chats": chats,
	})
}

func (h *Handler) createChat(c *gin.Context) {

	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "createChat.html", gin.H{})
	}

	if c.Request.Method == "POST" {

		name := c.Request.FormValue("name")

		token, err := c.Cookie("AUTH")
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		userId, err := h.services.ParseToken(token)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		chatId, err := h.services.Chat.Create(name, userId)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		//Creating new Online Room for websocket connection
		err = h.services.Websocket.CreateRoom(chatId, name)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Redirect(302, "/")
	}
}

func (h *Handler) addUser(c *gin.Context) {

	chatId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid chatId")
		return
	}

	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "addUser.html", gin.H{})
	}

	if c.Request.Method == "POST" {

		username := c.Request.FormValue("username")

		err = h.services.Chat.AddUser(chatId, username)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.Redirect(302, "/chat/"+strconv.Itoa(chatId))
	}
}

func (h *Handler) joinRoom(c *gin.Context) {

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

	c.HTML(http.StatusOK, "chat.html", gin.H{
		"chats":    chats,
		"messages": messages,
		"chatId":   chatId,
		"adminId":  adminId,
		"userId":   userId,
	})
}
