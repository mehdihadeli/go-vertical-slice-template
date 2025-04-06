package queries

import (
	"context"
	"fmt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/mappings"
	customErrors "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/httperrors/customerrors"
	uuid "github.com/satori/go.uuid"
	"net/http"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/dtos"
)

type GetProductByIdQueryHandler struct {
	productRepository contracts.ProductRepository
}

func NewGetProductByIdHandler(productRepository contracts.ProductRepository) *GetProductByIdQueryHandler {
	return &GetProductByIdQueryHandler{productRepository: productRepository}
}

func (q *GetProductByIdQueryHandler) Handle(ctx context.Context, query *GetProductByIdQuery) (*dtos.GetProductByIdQueryResponse, error) {
	if query == nil {
		return nil, customErrors.NewApplicationErrorWithCode("query cannot be nil", http.StatusBadRequest)
	}

	if query.ProductID == uuid.Nil {
		return nil, customErrors.NewApplicationErrorWithCode("product id cannot be empty", http.StatusBadRequest)
	}

	product, err := q.productRepository.GetProductById(ctx, query.ProductID)

	if err != nil || product == nil {
		return nil, customErrors.NewApplicationErrorWrapWithCode(err, http.StatusNotFound, fmt.Sprintf("product with id %s not found", query.ProductID))
	}

	productDto := mappings.MapProductToProductDto(product)

	return &dtos.GetProductByIdQueryResponse{Product: productDto}, nil
}
