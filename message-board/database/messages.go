package database

import (
	"database/sql"
	"errors"
	"log"
	"message-board/msg"

	"github.com/google/uuid"
)

type MessageService struct {
	db         *sql.DB
	insertStmt *sql.Stmt
	queryStmt  *sql.Stmt
	updateStmt *sql.Stmt
	deleteStmt *sql.Stmt
}

var ErrNotFound = errors.New("not found")

func NewMessageService(db *sql.DB) *MessageService {
	createMessage := "INSERT INTO messages(id, text) VALUES (?, ?)"
	insertStmt, err := db.Prepare(createMessage)
	check(err)

	queryMessage := "SELECT id, text FROM messages"
	queryStmt, err := db.Prepare(queryMessage)
	check(err)

	updateMessage := "UPDATE messages SET text=? WHERE id=?"
	updateStmt, err := db.Prepare(updateMessage)
	check(err)

	deleteMessage := "DELETE FROM messages WHERE id=?"
	deleteStmt, err := db.Prepare(deleteMessage)

	return &MessageService{
		db:         db,
		insertStmt: insertStmt,
		queryStmt:  queryStmt,
		updateStmt: updateStmt,
		deleteStmt: deleteStmt,
	}
}

func (service *MessageService) InsertMessage(text string) msg.Message {
	id := uuid.New()
	log.Printf("Adding message %s to database.\n", id)

	_, err := service.insertStmt.Exec(id, text)
	check(err)

	return msg.Message{Id: id, Text: text}
}

func (service *MessageService) QueryMessages() []msg.Message {
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

func (service *MessageService) QueryMessage(id uuid.UUID) msg.Message {
	queryRow := service.db.QueryRow("SELECT id, text FROM messages WHERE id=?", id)
	var message msg.Message
	queryRow.Scan(&message.Id, &message.Text)
	return message
}

func (service *MessageService) UpdateMessage(id uuid.UUID, text string) (msg.Message, error) {
	message := service.QueryMessage(id)
	if message.Id == uuid.Nil {
		return msg.Message{}, ErrNotFound
	}

	_, err := service.updateStmt.Exec(text, id)
	check(err)

	return msg.Message{Id: id, Text: text}, nil
}

func (service *MessageService) DeleteMessage(id uuid.UUID) error {
	message := service.QueryMessage(id)
	if message.Id == uuid.Nil {
		return ErrNotFound
	}

	_, err := service.deleteStmt.Exec(id)
	check(err)
	return nil
}
