package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"message-board/database"
	"message-board/msg"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostMessageHandler(t *testing.T) {
	db := database.InitSqliteMem()
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
	db := database.InitSqliteMem()
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

func TestPostMessageHandlerConcurrent(t *testing.T) {
	db := database.InitSqliteMem()
	service := database.NewMessageService(db)
	handler := PostMessagesHandler{Service: service}

	r := gin.Default()
	r.POST("/messages", handler.Handler)

	messageCount := 5
	responses := []*httptest.ResponseRecorder{}
	for i := 0; i < messageCount; i++ {
		w := httptest.NewRecorder()
		responses = append(responses, w)
	}
	var wg sync.WaitGroup
	for i, writer := range responses {
		wg.Add(1)
		w := writer
		msgNum := i
		go func() {
			defer wg.Done()
			log.Println("posting message...")
			messageText := fmt.Sprintf("concurrent %d", msgNum)
			postMessage(messageText, r, w)
			log.Println("done posting.")
		}()
	}
	wg.Wait()

	for i, w := range responses {
		assert.Equal(t, http.StatusCreated, w.Code)

		var message msg.Message
		json.Unmarshal(w.Body.Bytes(), &message)
		messageText := fmt.Sprintf("concurrent %d", i)
		assert.Equal(t, message.Text, messageText)
	}
}

func postMessage(text string, r *gin.Engine, w http.ResponseWriter) {
	body := gin.H{
		"text": text,
	}
	bodyBytes, err := json.Marshal(body)
	check(err)
	bodyReader := bytes.NewReader(bodyBytes)
	req, _ := http.NewRequest("POST", "/messages", bodyReader)
	r.ServeHTTP(w, req)
}
