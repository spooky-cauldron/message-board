package handlers

import (
	"database/sql"
	"message-board/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostMessagesHandler struct {
	Db *sql.DB
}

type PostMessageBody struct {
	Text string `json:"text" binding:"required"`
}

func (handler *PostMessagesHandler) Handler(c *gin.Context) {
	var body PostMessageBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := database.InsertMessage(handler.Db, body.Text)
	c.JSON(http.StatusCreated, message)
}
