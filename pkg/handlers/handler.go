package handlers

import (
	"mini_chat_go/pkg/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitHandlers() *gin.Engine {
	router := gin.New()
	auth := router.Group("/")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		tweets := api.Group("/tweets")
		{
			tweets.POST("", h.createTweet)
			tweets.GET("", h.getAllTweets)
			tweets.GET("/:id", h.getTweetById)
			tweets.DELETE("/:id", h.deleteTweet)
			tweets.PUT("/:id", h.updateTweet)
			// }
		}

		return router
	}
}
