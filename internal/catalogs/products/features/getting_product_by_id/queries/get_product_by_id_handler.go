package queries

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/getting_product_by_id/dtos"
)

type GetProductByIdQueryHandler struct {
	productRepository contracts.ProductRepository
}

func NewGetProductByIdHandler(productRepository contracts.ProductRepository) *GetProductByIdQueryHandler {
	return &GetProductByIdQueryHandler{productRepository: productRepository}
}

func (q *GetProductByIdQueryHandler) Handle(ctx context.Context, query *GetProductByIdQuery) (*dtos.GetProductByIdQueryResponse, error) {
	product, err := q.productRepository.GetProductById(ctx, query.ProductID)

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("product with id %s not found", query.ProductID))
	}

	productDto := products.MapProductToProductDto(product)

	return &dtos.GetProductByIdQueryResponse{Product: productDto}, nil
}
