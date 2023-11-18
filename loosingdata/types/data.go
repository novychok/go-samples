package types

import "github.com/google/uuid"

type Data struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

func NewData(text string) *Data {
	return &Data{
		ID:   uuid.New(),
		Text: text,
	}
}
