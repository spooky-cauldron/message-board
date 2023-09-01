package database

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertMessage(t *testing.T) {
	db := InitSqliteMem()
	service := NewMessageService(db)
	testMessageText := "testing message"
	inserted := service.InsertMessage(testMessageText)

	messageText := inserted.Text
	assert.Equal(t, messageText, testMessageText)
}

func TestQueryMessages(t *testing.T) {
	db := InitSqliteMem()
	service := NewMessageService(db)
	testMessageText := "testing message"
	service.InsertMessage(testMessageText)

	messages := service.QueryMessages()
	assert.Len(t, messages, 1)

	message := messages[0]
	assert.Equal(t, message.Text, testMessageText)

	assert.NotEqual(t, message.Id, uuid.Nil)
}
