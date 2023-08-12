package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) homePage(c *gin.Context) {

	token := ""
	if cookie, err := c.Request.Cookie("AUTH"); err == nil {
		token = cookie.Value
	}
	log.Print(token)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Chat it all",
		"token": token,
	})
}
