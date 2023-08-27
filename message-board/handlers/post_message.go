package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"message-board/msg"
	"net/http"

	"github.com/google/uuid"
)

type PostMessagesHandler struct {
	Db *sql.DB
}

type PostMessageBody struct {
	Text string `json:"text"`
}

func (handler *PostMessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := AssertPost(w, r); err != nil {
		return
	}

	bodyData, err := io.ReadAll(r.Body)
	check(err)

	var body PostMessageBody
	err = json.Unmarshal(bodyData, &body)
	check(err)

	message := InsertMessage(handler.Db, body.Text)
	WriteJson(message, w)
}

func InsertMessage(db *sql.DB, text string) msg.Message {
	createMessage := "INSERT INTO messages(id, text) VALUES (?, ?)"
	stmt, err := db.Prepare(createMessage)
	check(err)

	id := uuid.New()
	log.Printf("Adding message %s to database.\n", id)

	res, err := stmt.Exec(id, text)
	check(err)
	fmt.Println(res.LastInsertId())

	return msg.Message{
		Id:   id,
		Text: text,
	}
}
