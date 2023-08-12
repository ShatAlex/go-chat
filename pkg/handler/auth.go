package handler

import (
	"net/http"

	"github.com/ShatAlex/chat"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	}
	if c.Request.Method == "POST" {

		name := c.Request.FormValue("name")
		username := c.Request.FormValue("username")
		password1 := c.Request.FormValue("password1")
		password2 := c.Request.FormValue("password2")

		if err := h.validatePassword(password1, password2); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		input := chat.User{
			Name:     name,
			Username: username,
			Password: password1,
		}

		_, err := h.services.CreateUser(input)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.SetCookie("AUTH", token, 3600, "/", "127.0.0.1", false, true)

		c.Redirect(302, "/")
	}

}

func (h *Handler) signIn(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}
	if c.Request.Method == "POST" {

		username := c.Request.FormValue("username")
		password := c.Request.FormValue("password")

		input := chat.SignInUser{
			Username: username,
			Password: password,
		}

		token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.SetCookie("AUTH", token, 3600, "/", "127.0.0.1", false, true)

		c.Redirect(302, "/")

	}

}

func (h *Handler) signOut(c *gin.Context) {

	c.SetCookie("AUTH", "", 1, "/", "127.0.0.1", false, true)

	c.Redirect(302, "/")

}
