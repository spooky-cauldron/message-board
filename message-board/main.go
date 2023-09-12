package main

import (
	"database/sql"
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

	db := initDb()
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

	host := getEnv("HOST", "localhost:8000")

	err := router.Run(host)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed.")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}

func initDb() *sql.DB {
	var db *sql.DB
	switch getEnv("DB_TYPE", "postgres") {
	case "sqlite-memory":
		log.Println("Using SQLite in-memory database.")
		db = database.InitSqliteMem()
	case "sqlite":
		dbFile := getEnv("SQLITE_FILE", "messageboard.db")
		log.Printf("Using SQLite database at file %s\n", dbFile)
		db = database.InitSqlite(dbFile)
	default:
		log.Println("Using Postgres database.")
		db = database.InitPostgres()
	}
	return db
}

func getEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
