package handlers

import (
	"bytes"
	"encoding/json"
	"message-board/database"
	"message-board/msg"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostMessageHandler(t *testing.T) {
	db := database.InitSqlite()
	service := database.NewMessageService(db)
	handler := PostMessagesHandler{Service: service}

	r := gin.Default()
	r.POST("/messages", handler.Handler)

	body := gin.H{
		"text": "test post",
	}
	bodyBytes, err := json.Marshal(body)
	check(err)
	bodyReader := bytes.NewReader(bodyBytes)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/messages", bodyReader)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var message msg.Message
	json.Unmarshal(w.Body.Bytes(), &message)
	assert.Equal(t, message.Text, "test post")
}

func TestPostMessageHandlerMissingData(t *testing.T) {
	db := database.InitSqlite()
	service := database.NewMessageService(db)
	handler := PostMessagesHandler{Service: service}

	r := gin.Default()
	r.POST("/messages", handler.Handler)

	body := gin.H{}
	bodyBytes, err := json.Marshal(body)
	check(err)
	bodyReader := bytes.NewReader(bodyBytes)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/messages", bodyReader)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "required")
}
