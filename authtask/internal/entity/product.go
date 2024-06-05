package entity

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")
)

type Product struct {
	ID                 int    `json:"id"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductPrice       string `json:"product_price"`
}
