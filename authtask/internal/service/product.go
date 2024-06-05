package service

import (
	"context"

	"github.com/novychok/go-samples/authtask/internal/entity"
)

type ProductService interface {
	GetByName(ctx context.Context, productName string) (*entity.Product, error)
}
