package queries

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/go-vertical-slice-template/internal/products"
	"github.com/go-vertical-slice-template/internal/products/contracts"
	gettingProductByIdDtos "github.com/go-vertical-slice-template/internal/products/features/getting_product_by_id/dtos"
)

type GetProductByIdQueryHandler struct {
	productRepository contracts.ProductRepository
}

func NewGetProductByIdHandler(productRepository contracts.ProductRepository) *GetProductByIdQueryHandler {
	return &GetProductByIdQueryHandler{productRepository: productRepository}
}

func (q *GetProductByIdQueryHandler) Handle(ctx context.Context, query *GetProductByIdQuery) (*gettingProductByIdDtos.GetProductByIdQueryResponse, error) {
	product, err := q.productRepository.GetProductById(ctx, query.ProductID)

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("product with id %s not found", query.ProductID))
	}

	productDto := products.MapProductToProductDto(product)

	return &gettingProductByIdDtos.GetProductByIdQueryResponse{Product: productDto}, nil
}
