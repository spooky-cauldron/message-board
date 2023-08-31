package handlers

import (
	"message-board/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMessagesHandler struct {
	Service *database.MessageService
}

func (h *GetMessagesHandler) Handler(c *gin.Context) {
	messages := h.Service.QueryMessages()
	c.JSON(http.StatusOK, messages)
}
