package handlers

import (
	"net/http"
	"strconv"
	tablepractise "mini_chat_go"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTweet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input tablepractise.TodoTweet
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := h.service.TodoTweet.CreateTweet(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllTweetsResponce struct {
	Data []tablepractise.TodoTweet `json:"data"`
}

func (h *Handler) getAllTweets(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	tweets, err := h.service.TodoTweet.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTweetsResponce{
		Data: tweets,
	})
}

func (h *Handler) getTweetById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	tweet, err := h.service.TodoTweet.GetTweetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tweet)
}

func (h *Handler) deleteTweet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id content param")
		return
	}
	err = h.service.TodoTweet.DeleteTweet(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}
func (h *Handler) updateTweet(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	return
	// }
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Method Under Developing",
	})
}
