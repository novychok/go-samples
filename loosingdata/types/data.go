package types

import "github.com/google/uuid"

type Data struct {
	ID          uuid.UUID `json:"id"`
	Message     string    `json:"message"`
	MessageHash string    `json:"message_hash"`
}

func NewData(msg, msgHash string) *Data {
	return &Data{
		ID:          uuid.New(),
		Message:     msg,
		MessageHash: msgHash,
	}
}
