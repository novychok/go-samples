package repository

import (
	"context"

	"github.com/novychok/go-samples/authtask/internal/entity"
)

type ProductRepository interface {
	GetByName(ctx context.Context, productName string) (*entity.Product, error)
}
