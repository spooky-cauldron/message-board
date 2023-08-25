package msg

import "github.com/google/uuid"

type Message struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
