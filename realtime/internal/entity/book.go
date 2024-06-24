package entity

import "errors"

var (
	ErrBookAlreadyExists = errors.New("error: book already exists")
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
