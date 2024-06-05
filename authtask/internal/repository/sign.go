package repository

import (
	"context"

	"github.com/novychok/go-samples/authtask/internal/entity"
)

type SignRepository interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, username string) (*entity.User, error)
}
