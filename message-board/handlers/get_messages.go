package handlers

import (
	"database/sql"
	"message-board/msg"
	"net/http"
)

type GetMessagesHandler struct {
	Db *sql.DB
}

func (h *GetMessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := AssertGet(w, r); err != nil {
		return
	}

	messages := queryMessages(h.Db)
	WriteJson(messages, w)
}

func queryMessages(db *sql.DB) []msg.Message {
	rows, err := db.Query("SELECT id, text FROM messages")
	check(err)
	defer rows.Close()

	var messages []msg.Message
	for rows.Next() {
		var message msg.Message
		err := rows.Scan(&message.Id, &message.Text)
		check(err)
		messages = append(messages, message)
	}
	return messages
}

func queryMessage(db *sql.DB) msg.Message {
	queryRow := db.QueryRow("SELECT id, text FROM messages")
	var message msg.Message
	queryRow.Scan(&message.Id, &message.Text)
	return message
}
