package handlers

import (
	"message-board/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PatchMessageHandler struct {
	Service *database.MessageService
}

type PatchMessageBody struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Text string    `json:"text" binding:"required"`
}

func (h *PatchMessageHandler) Handler(c *gin.Context) {
	var body PatchMessageBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.Service.UpdateMessage(body.Id, body.Text)
	if err == database.ErrNotFound {
		http.NotFound(c.Writer, c.Request)
		return
	}
	c.JSON(http.StatusOK, message)
}
