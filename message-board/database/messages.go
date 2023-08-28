package database

import (
	"database/sql"
	"log"
	"message-board/msg"

	"github.com/google/uuid"
)

func InsertMessage(db *sql.DB, text string) msg.Message {
	createMessage := "INSERT INTO messages(id, text) VALUES (?, ?)"
	id := uuid.New()
	log.Printf("Adding message %s to database.\n", id)

	_, err := db.Exec(createMessage, id, text)
	check(err)

	return msg.Message{Id: id, Text: text}
}

func QueryMessages(db *sql.DB) []msg.Message {
	rows, err := db.Query("SELECT id, text FROM messages")
	check(err)
	defer rows.Close()

	messages := []msg.Message{}
	for rows.Next() {
		var message msg.Message
		err := rows.Scan(&message.Id, &message.Text)
		check(err)
		messages = append(messages, message)
	}
	return messages
}

func QueryMessage(db *sql.DB) msg.Message {
	queryRow := db.QueryRow("SELECT id, text FROM messages")
	var message msg.Message
	queryRow.Scan(&message.Id, &message.Text)
	return message
}
