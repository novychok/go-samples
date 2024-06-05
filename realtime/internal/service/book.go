package service

import (
	"context"

	"github.com/novychok/go-samples/realtime/internal/entity"
)

type Books interface {
	FindAll(ctx context.Context) ([]*entity.Book, error)
	Create(ctx context.Context, book *entity.Book) (*entity.Book, error)
	GetByID(ctx context.Context, id string) (*entity.Book, error)
}
