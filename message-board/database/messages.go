package database

import (
	"database/sql"
	"log"
	"message-board/msg"

	"github.com/google/uuid"
)

type MessageService struct {
	db         *sql.DB
	insertStmt *sql.Stmt
	queryStmt  *sql.Stmt
}

func NewMessageService(db *sql.DB) *MessageService {
	createMessage := "INSERT INTO messages(id, text) VALUES (?, ?)"
	insertStmt, err := db.Prepare(createMessage)
	check(err)

	queryMessage := "SELECT id, text FROM messages"
	queryStmt, err := db.Prepare(queryMessage)
	check(err)

	return &MessageService{db: db, insertStmt: insertStmt, queryStmt: queryStmt}
}

func (service *MessageService) InsertMessage(text string) msg.Message {
	id := uuid.New()
	log.Printf("Adding message %s to database.\n", id)

	// _, err := service.db.Exec("INSERT INTO messages(id, text) VALUES (?, ?)", id, text)
	_, err := service.insertStmt.Exec(id, text)
	check(err)

	return msg.Message{Id: id, Text: text}
}

func (service *MessageService) QueryMessages() []msg.Message {
	// rows, err := service.db.Query("SELECT id, text FROM messages")
	rows, err := service.queryStmt.Query()
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
