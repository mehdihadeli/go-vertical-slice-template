package dtos

import (
	"github.com/go-vertical-slice-template/internal/products/dtos"
)

type GetProductByIdQueryResponse struct {
	Product *dtos.ProductDto `json:"product"`
}
