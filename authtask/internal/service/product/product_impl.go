package product

import (
	"context"
	"log/slog"

	"github.com/novychok/go-samples/authtask/internal/entity"
	"github.com/novychok/go-samples/authtask/internal/repository"
	"github.com/novychok/go-samples/authtask/internal/service"
)

type srv struct {
	productRepository repository.ProductRepository
	l                 *slog.Logger
}

func (s *srv) GetByName(ctx context.Context, productName string) (*entity.Product, error) {
	l := s.l.With(slog.String("method", "getByName"))

	product, err := s.productRepository.GetByName(ctx, productName)
	if err != nil {
		l.ErrorContext(ctx, "Failed to get product from database", "err", err)
		return nil, err
	}

	return product, nil
}

func New(productRepository repository.ProductRepository, l *slog.Logger) service.ProductService {
	return &srv{
		productRepository: productRepository,
		l:                 l,
	}
}
