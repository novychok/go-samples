package product

import (
	"context"
	"database/sql"

	"github.com/novychok/go-samples/authtask/internal/entity"
	"github.com/novychok/go-samples/authtask/internal/repository"
)

type psql struct {
	db *sql.DB
}

func (r *psql) GetByName(ctx context.Context, productName string) (*entity.Product, error) {
	var product entity.Product
	row := r.db.QueryRow("SELECT * FROM products WHERE product_name = $1", productName)

	err := row.Scan(&product.ID, &product.ProductName,
		&product.ProductDescription, &product.ProductPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil
}

func New(db *sql.DB) repository.ProductRepository {
	return &psql{db: db}
}
