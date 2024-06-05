package sign

import (
	"context"
	"database/sql"

	"github.com/novychok/go-samples/authtask/internal/entity"
	"github.com/novychok/go-samples/authtask/internal/repository"

	"github.com/lib/pq"
)

type psql struct {
	db *sql.DB
}

var (
	constraintErrorCode = "23505"
)

func (r *psql) Create(ctx context.Context, user *entity.User) error {
	query := "INSERT INTO users (username, password_hash) values ($1, $2) RETURNING id"

	_, err := r.db.Exec(query, user.Username, user.PasswordHash)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == pq.ErrorCode(constraintErrorCode) {
			return entity.ErrUserAlreadyExists
		}
	}

	return nil
}

func (r *psql) Get(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	row := r.db.QueryRow("SELECT * FROM users WHERE username = $1", username)

	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func New(db *sql.DB) repository.SignRepository {
	return &psql{db: db}
}
