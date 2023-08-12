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

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("./pkg/templates/*")
	router.Static("/static", "./pkg/static/")

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // expire in a day
	router.Use(sessions.Sessions("auth", store))

	router.GET("/", h.homePage)

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", h.signUp)
		auth.POST("/sign-up", h.signUp)

		auth.GET("/sign-in", h.signIn)
		auth.POST("/sign-in", h.signIn)

		auth.GET("/sign-out", h.signOut)
	}

	return router
}
