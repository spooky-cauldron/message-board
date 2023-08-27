package msg

import "github.com/google/uuid"

type Message struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}
