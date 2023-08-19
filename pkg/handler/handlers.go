package handler

import (
	"github.com/ShatAlex/chat/pkg/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(ser *service.Service) *Handler {
	return &Handler{services: ser}
}

func (h *Handler) InitRouters(wsh *WsHandler) *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("./pkg/templates/*")
	router.Static("/static", "./pkg/static/")

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // expire in a day
	router.Use(sessions.Sessions("auth", store))

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", h.signUp)
		auth.POST("/sign-up", h.signUp)

		auth.GET("/sign-in", h.signIn)
		auth.POST("/sign-in", h.signIn)

		auth.GET("/sign-out", h.signOut)
	}

	router.GET("/", h.homePage)
	router.GET("/create-chat", h.createChat)
	router.POST("/create-chat", h.createChat)

	router.GET("chat/:id", wsh.joinRoom)

	router.GET("chat/:id/add-user", h.addUser)
	router.POST("chat/:id/add-user", h.addUser)

	return router
}
