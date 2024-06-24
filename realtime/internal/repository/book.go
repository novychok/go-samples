package repository

import (
	"context"

	"github.com/novychok/go-samples/realtime/internal/entity"
)

type Book interface {
	Create(ctx context.Context, create *entity.Book) (int, error)
}
