package main

import (
	"errors"
	"log"
	"message-board/database"
	"message-board/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitSqliteMem()
	messageService := database.NewMessageService(db)

	r := gin.Default()

	getMessageHandler := handlers.GetMessagesHandler{Service: messageService}
	r.GET("/message", getMessageHandler.Handler)

	postMessageHandler := handlers.PostMessagesHandler{Service: messageService}
	r.POST("/message", postMessageHandler.Handler)

	patchMessageHandler := handlers.PatchMessageHandler{Service: messageService}
	r.PATCH("/message", patchMessageHandler.Handler)

	deleteMessageHandler := handlers.DeleteMessageHandler{Service: messageService}
	r.DELETE("/message/:id", deleteMessageHandler.Handler)

	err := r.Run("localhost:8000")

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed.")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
