package main

import (
	"errors"
	"log"
	"message-board/database"
	"message-board/handlers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Initializing Service...")
	db := database.InitSqliteMem()
	messageService := database.NewMessageService(db)

	router := gin.Default()

	getMessageHandler := handlers.GetMessagesHandler{Service: messageService}
	router.GET("/message", getMessageHandler.Handler)

	postMessageHandler := handlers.PostMessagesHandler{Service: messageService}
	router.POST("/message", postMessageHandler.Handler)

	patchMessageHandler := handlers.PatchMessageHandler{Service: messageService}
	router.PATCH("/message", patchMessageHandler.Handler)

	deleteMessageHandler := handlers.DeleteMessageHandler{Service: messageService}
	router.DELETE("/message/:id", deleteMessageHandler.Handler)

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = "localhost:8000"
	}

	err := router.Run(host)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed.")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
