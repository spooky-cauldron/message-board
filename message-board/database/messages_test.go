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

func TestUpdateMessage(t *testing.T) {
	db := InitSqliteMem()
	service := NewMessageService(db)
	initialText := "initial message"
	message := service.InsertMessage(initialText)

	assert.Equal(t, initialText, message.Text)

	updatedText := "updated"
	updatedMessage, err := service.UpdateMessage(message.Id, updatedText)

	assert.Equal(t, nil, err)
	assert.Equal(t, updatedText, updatedMessage.Text)
	assert.NotEqual(t, message.Text, updatedMessage.Text)
}

func TestUpdateMessageNotFound(t *testing.T) {
	db := InitSqliteMem()
	service := NewMessageService(db)

	updatedText := "updated"
	updatedMessage, err := service.UpdateMessage(uuid.New(), updatedText)

	assert.Equal(t, ErrNotFound, err)
	assert.NotEqual(t, updatedText, updatedMessage.Text)
	assert.Equal(t, uuid.Nil, updatedMessage.Id)
}

func TestDeleteMessage(t *testing.T) {
	db := InitSqliteMem()
	service := NewMessageService(db)
	initialText := "delete message"
	message := service.InsertMessage(initialText)

	assert.Equal(t, initialText, message.Text)

	err := service.DeleteMessage(message.Id)
	assert.Equal(t, nil, err)

	deleted := service.QueryMessage(message.Id)
	assert.Equal(t, uuid.Nil, deleted.Id)
}

func TestDeleteMessageNotFound(t *testing.T) {
	db := InitSqliteMem()
	service := NewMessageService(db)

	err := service.DeleteMessage(uuid.New())
	assert.Equal(t, ErrNotFound, err)
}
