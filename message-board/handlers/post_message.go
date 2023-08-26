package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"message-board/msg"
	"net/http"

	"github.com/google/uuid"
)

type PostMessagesHandler struct {
	Db *sql.DB
}

type PostMessageBody struct {
	Name string `json:"name"`
}

func (h *PostMessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := AssertPost(w, r); err != nil {
		return
	}

	bodyData, err := io.ReadAll(r.Body)
	check(err)

	var body PostMessageBody
	err = json.Unmarshal(bodyData, &body)
	check(err)

	message := InsertMessage(h.Db, body.Name)
	WriteJson(message, w)
}

func InsertMessage(db *sql.DB, name string) msg.Message {
	createMessage := "INSERT INTO messages(id, name) VALUES (?, ?)"
	stmt, err := db.Prepare(createMessage)
	check(err)

	id := uuid.New()
	fmt.Println("id")
	fmt.Println(id)

	res, err := stmt.Exec(id, name)
	check(err)
	fmt.Println(res.LastInsertId())

	return msg.Message{
		Id:   id,
		Name: name,
	}
}
