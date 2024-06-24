package book

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/novychok/go-samples/realtime/internal/domain"
	"github.com/novychok/go-samples/realtime/internal/entity"
	"github.com/novychok/go-samples/realtime/internal/pkg/postgres"
	"github.com/novychok/go-samples/realtime/internal/repository"
)

var (
	constraintErrorCode = "23505"
)

const (
	EVENT_CREATE = "books.create"
)

type psql struct {
	db *sql.DB
}

// TODO: DAO structure
type event struct {
	ID      int    `db:"id"`
	Type    string `db:"event"`
	Payload string `db:"payload"`
}

func (r *psql) Create(ctx context.Context, create *entity.Book) (id int, err error) {

	query := `INSERT INTO books (title, name, author) values ($1, $2, $3) RETURNING id`

	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("begin tx: %s", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}

		if err := tx.Commit(); err != nil {
			log.Printf("error to commit transaction: %s", err)
			return
		}
	}()

	err = tx.QueryRowContext(ctx, query, create.Title, create.Name, create.Author).Scan(&id)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == pq.ErrorCode(constraintErrorCode) {
			return 0, entity.ErrBookAlreadyExists
		}
		return 0, fmt.Errorf("error query row ctx: %s", err)
	}

	payload := fmt.Sprintf(`{"id": %d}`, id)
	if err := r.saveEvent(ctx, tx, EVENT_CREATE, payload); err != nil {
		return 0, fmt.Errorf("error save event: %s", err)
	}

	return id, nil
}

func (r *psql) saveEvent(ctx context.Context, tx *sql.Tx, eventType, payload string) error {

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO events (event_type, payload) VALUES($1, $2)")
	if err != nil {
		return fmt.Errorf("prepare save tx: %s", err)
	}

	_, err = stmt.Exec(eventType, payload)
	if err != nil {
		return fmt.Errorf("exec save tx: %s", err)
	}

	return nil
}

func (r *psql) GetEvent(ctx context.Context) (*domain.Event, error) {

	row := r.db.QueryRowContext(ctx, "SELECT * FROM events WHERE status = 'new' LIMIT 1")

	var evt event
	err := row.Scan(&evt.ID, &evt.Type, &evt.Payload)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no rows specfied: %s", err)
		}
		return nil, fmt.Errorf("err scan: %s", err)
	}

	return &domain.Event{
		ID:      evt.ID,
		Type:    evt.Type,
		Payload: evt.Payload,
	}, nil
}

func (r *psql) SetDone(ctx context.Context, id int) error {

	stmt, err := r.db.PrepareContext(ctx, "UPDATE events SET status = 'done' WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func NewPostgres(connection postgres.Connection) repository.Book {
	return &psql{
		db: connection,
	}
}
