package repository

import (
	"context"

	"github.com/novychok/go-samples/realtime/internal/domain"
	"github.com/novychok/go-samples/realtime/internal/entity"
)

type Book interface {
	Create(ctx context.Context, create *entity.Book) (int, error)
	GetEvent(ctx context.Context) (*domain.Event, error)
	SetDone(ctx context.Context, id int) error
}
