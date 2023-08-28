package handlers

import (
	"database/sql"
	"message-board/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMessagesHandler struct {
	Db *sql.DB
}

func (h *GetMessagesHandler) Handler(c *gin.Context) {
	messages := database.QueryMessages(h.Db)
	c.JSON(http.StatusOK, messages)
}
