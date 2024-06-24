package service

import (
	"context"

	"github.com/novychok/go-samples/realtime/internal/entity"
)

type Books interface {
	Create(ctx context.Context, book *entity.Book) (int, error)
}
