package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"message-board/msg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostMessagesHandler struct {
	Db *sql.DB
}

type PostMessageBody struct {
	Text string `json:"text"`
}

func (handler *PostMessagesHandler) Handler(c *gin.Context) {
	bodyData, err := io.ReadAll(c.Request.Body)
	check(err)

	var body PostMessageBody
	err = json.Unmarshal(bodyData, &body)
	check(err)

	message := InsertMessage(handler.Db, body.Text)
	c.JSON(http.StatusOK, message)
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
