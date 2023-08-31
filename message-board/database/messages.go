package database

import (
	"database/sql"
	"log"
	"message-board/msg"

	"github.com/google/uuid"
)

type MessageService struct {
	db *sql.DB
}

func NewMessageService(db *sql.DB) *MessageService {
	return &MessageService{db: db}
}

func (service *MessageService) InsertMessage(text string) msg.Message {
	createMessage := "INSERT INTO messages(id, text) VALUES (?, ?)"
	id := uuid.New()
	log.Printf("Adding message %s to database.\n", id)

	_, err := service.db.Exec(createMessage, id, text)
	check(err)

	return msg.Message{Id: id, Text: text}
}

func (service *MessageService) QueryMessages() []msg.Message {
	rows, err := service.db.Query("SELECT id, text FROM messages")
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

func (service *MessageService) QueryMessage() msg.Message {
	queryRow := service.db.QueryRow("SELECT id, text FROM messages")
	var message msg.Message
	queryRow.Scan(&message.Id, &message.Text)
	return message
}
