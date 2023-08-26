package main

import (
	"errors"
	"fmt"
	"log"
	"message-board/database"
	"message-board/handlers"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := database.InitSqlite()

	handlers.InsertMessage(db, "hi")

	http.Handle("/message", &handlers.GetMessagesHandler{Db: db})
	http.Handle("/new-message", &handlers.PostMessagesHandler{Db: db})

	fmt.Println("Starting server.")
	err := http.ListenAndServe(":8000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed.")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
