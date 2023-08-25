package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"message-board/handlers"
	"net/http"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDb()

	insertMessage(db)

	http.Handle("/message", &handlers.GetMessageHandler{Db: db})

	fmt.Println("Starting server.")
	err := http.ListenAndServe(":8000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed.")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func initDb() *sql.DB {
	dbSource := ":memory:"
	db, err := sql.Open("sqlite3", dbSource)
	check(err)

	createTableQuery, err := os.ReadFile("../sql/messages_table.sql")
	check(err)

	_, err = db.Exec(string(createTableQuery))
	check(err)

	log.Println("Initialized Database.")
	return db
}

func insertMessage(db *sql.DB) {
	createMessage := "INSERT INTO messages(id, name) VALUES (?, ?)"
	stmt, err := db.Prepare(createMessage)
	check(err)

	id := uuid.New()
	fmt.Println("id")
	fmt.Println(id)

	res, err := stmt.Exec(id, "hi")
	check(err)
	fmt.Println(res.LastInsertId())
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
