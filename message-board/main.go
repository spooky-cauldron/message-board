package main

import (
	"errors"
	"fmt"
	"log"
	"message-board/database"
	"message-board/handlers"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := database.InitSqlite()

	handlers.InsertMessage(db, "hi")

	r := gin.Default()

	getMessageHandler := handlers.GetMessagesHandler{Db: db}
	r.GET("/message", getMessageHandler.Handler)

	postMessageHandler := handlers.PostMessagesHandler{Db: db}
	r.POST("/message", postMessageHandler.Handler)

	err := r.Run("localhost:8000")

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed.")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}

// func check(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
