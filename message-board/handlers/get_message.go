package handlers

import (
	"database/sql"
	"message-board/msg"
	"net/http"
)

type GetMessageHandler struct {
	Db *sql.DB
}

func (h *GetMessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := queryMessage(h.Db)
	WriteJson(message, w)
}

func queryMessage(db *sql.DB) msg.Message {
	queryRow := db.QueryRow("SELECT id, name FROM messages")
	var message msg.Message
	queryRow.Scan(&message.Id, &message.Name)
	return message
}
