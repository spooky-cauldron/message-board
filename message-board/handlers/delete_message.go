package handlers

import (
	"message-board/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteMessageHandler struct {
	Service *database.MessageService
}

func (h *DeleteMessageHandler) Handler(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	err = h.Service.DeleteMessage(id)
	if err == database.ErrNotFound {
		http.NotFound(c.Writer, c.Request)
		return
	}
	c.Status(http.StatusNoContent)
}
