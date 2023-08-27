package handlers

import (
	"database/sql"
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
	Text string `json:"text" binding:"required"`
}

func (handler *PostMessagesHandler) Handler(c *gin.Context) {
	var body PostMessageBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := InsertMessage(handler.Db, body.Text)
	c.JSON(http.StatusOK, message)
}

func InsertMessage(db *sql.DB, text string) msg.Message {
	createMessage := "INSERT INTO messages(id, text) VALUES (?, ?)"
	stmt, err := db.Prepare(createMessage)
	check(err)

	id := uuid.New()
	log.Printf("Adding message %s to database.\n", id)

	_, err = stmt.Exec(id, text)
	check(err)

	return msg.Message{Id: id, Text: text}
}
