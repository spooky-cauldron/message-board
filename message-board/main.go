package main

import (
	"errors"
	"log"
	"message-board/database"
	"message-board/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := database.InitSqlite()
	database.InsertMessage(db, "test message")

	r := gin.Default()

	getMessageHandler := handlers.GetMessagesHandler{Db: db}
	r.GET("/message", getMessageHandler.Handler)

	postMessageHandler := handlers.PostMessagesHandler{Db: db}
	r.POST("/message", postMessageHandler.Handler)

	err := r.Run("localhost:8000")

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed.")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
