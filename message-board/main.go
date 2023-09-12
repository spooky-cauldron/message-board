package main

import (
	"database/sql"
	"errors"
	"fmt"
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
	db_host := getEnv("DB_HOST", "localhost")
	db_port := getEnv("DB_PORT", "5432")
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_ssl := getEnv("DB_SSL", "disable")
	connectionData := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		db_host,
		db_port,
		db_name,
		db_user,
		db_pass,
		db_ssl,
	)
	db := database.InitPostgres(connectionData)
	return db
}

func getEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
