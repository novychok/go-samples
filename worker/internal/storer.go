package internal

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/novychok/go-samples/worker/types"
)

type Storer struct {
	db *sql.DB
}

func NewStorer(db *sql.DB) *Storer {
	return &Storer{db: db}
}

func (s *Storer) StoreObject(object types.Object) error {
	_, err := s.db.Exec(`INSERT INTO object(id, online, last_seen)
		VALUES($1, $2, $3)`,
		object.ID, object.Online, time.Now())
	if err != nil {
		fmt.Printf("error insert in database %v\n", err)
		return nil
	}

	return nil
}

func (s *Storer) InitSchemas() error {
	query := `
	CREATE TABLE IF NOT EXISTS object(
		id INTEGER NOT NULL,
		online BOOLEAN NOT NULL,
		last_seen DATE NOT NULL
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *Storer) DropSchemas() error {
	query := `DROP TABLE object`
	_, err := s.db.Exec(query)
	return err
}
