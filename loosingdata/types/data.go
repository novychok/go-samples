package types

import "github.com/google/uuid"

type Data struct {
	ID       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	TextHash string    `json:"text_hash"`
}

func NewData(text, textHash string) *Data {
	return &Data{
		ID:       uuid.New(),
		Text:     text,
		TextHash: textHash,
	}
}
