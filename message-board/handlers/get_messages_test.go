package handlers

import (
	"encoding/json"
	"message-board/database"
	"message-board/msg"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetMessageHandler(t *testing.T) {
	db := database.InitSqliteMem()
	service := database.NewMessageService(db)
	service.InsertMessage("test message")
	handler := GetMessagesHandler{Service: service}

	r := gin.Default()
	r.GET("/messages", handler.Handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/messages", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var data []msg.Message
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Len(t, data, 1)
	message := data[0]
	assert.Equal(t, message.Text, "test message")
}

func TestGetMessageHandlerEmptyResponse(t *testing.T) {
	db := database.InitSqliteMem()
	service := database.NewMessageService(db)
	handler := GetMessagesHandler{Service: service}

	r := gin.Default()
	r.GET("/messages", handler.Handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/messages", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var data []msg.Message
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Len(t, data, 0)
}
