package book

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/novychok/go-samples/realtime/internal/entity"
	"github.com/novychok/go-samples/realtime/internal/pkg/postgres"
	"github.com/novychok/go-samples/realtime/internal/repository"
)

var (
	constraintErrorCode = "23505"
)

type psql struct {
	db *sql.DB
}

func (r *psql) Create(ctx context.Context, create *entity.Book) (int, error) {

	query := `INSERT INTO books (title, name, author) values ($1, $2, $3) RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query, create.Title, create.Name, create.Author).Scan(&id)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == pq.ErrorCode(constraintErrorCode) {
			return 0, entity.ErrBookAlreadyExists
		}
		return 0, err
	}

	return id, nil
}

func NewPostgres(connection postgres.Connection) repository.Book {
	return &psql{
		db: connection,
	}
}
