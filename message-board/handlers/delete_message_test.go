package handlers

import (
	"fmt"
	"message-board/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteMessageHandler(t *testing.T) {
	db := database.InitSqliteMem()
	service := database.NewMessageService(db)
	message := service.InsertMessage("delete test")
	handler := DeleteMessageHandler{Service: service}

	r := gin.Default()
	r.DELETE("/messages/:id", handler.Handler)

	w := httptest.NewRecorder()
	requestPath := fmt.Sprintf("/messages/%s", message.Id)
	req, _ := http.NewRequest(http.MethodDelete, requestPath, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteMessageHandlerNotFound(t *testing.T) {
	db := database.InitSqliteMem()
	service := database.NewMessageService(db)
	handler := DeleteMessageHandler{Service: service}

	r := gin.Default()
	r.DELETE("/messages/:id", handler.Handler)

	w := httptest.NewRecorder()
	requestPath := fmt.Sprintf("/messages/%s", uuid.New())
	req, _ := http.NewRequest(http.MethodDelete, requestPath, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
