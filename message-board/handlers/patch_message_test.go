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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPatchMessageHandler(t *testing.T) {
	db := database.InitSqliteMem()
	service := database.NewMessageService(db)
	message := service.InsertMessage("patch test")
	handler := PatchMessageHandler{Service: service}

	r := gin.Default()
	r.PATCH("/messages", handler.Handler)

	body := gin.H{
		"id":   message.Id,
		"text": "test post",
	}
	bodyBytes, err := json.Marshal(body)
	check(err)
	bodyReader := bytes.NewReader(bodyBytes)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/messages", bodyReader)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var editedMessage msg.Message
	json.Unmarshal(w.Body.Bytes(), &editedMessage)
	assert.Equal(t, editedMessage.Text, "test post")
}

func TestPatchMessageHandlerNotFound(t *testing.T) {
	db := database.InitSqliteMem()
	service := database.NewMessageService(db)
	handler := PatchMessageHandler{Service: service}

	r := gin.Default()
	r.PATCH("/messages", handler.Handler)

	body := gin.H{
		"id":   uuid.New(),
		"text": "test post",
	}
	bodyBytes, err := json.Marshal(body)
	check(err)
	bodyReader := bytes.NewReader(bodyBytes)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/messages", bodyReader)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
