package types

import "time"

type ObjectsIDs struct {
	ID []int `json:"object_ids"`
}

type Object struct {
	ID       int       `json:"id"`
	Online   bool      `json:"online"`
	LastSeen time.Time `json:"last_seen"`
}
