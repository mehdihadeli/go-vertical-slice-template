package dtos

import "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/dtos"

type GetProductByIdQueryResponse struct {
	Product *dtos.ProductDto `json:"product"`
}
