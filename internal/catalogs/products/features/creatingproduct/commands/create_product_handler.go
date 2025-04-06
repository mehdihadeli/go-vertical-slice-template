package commands

import (
	"context"
	"net/http"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/events"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/mappings"
	customErrors "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/httperrors/customerrors"

	"github.com/mehdihadeli/go-mediatr"
	uuid "github.com/satori/go.uuid"
)

type CreateProductCommandHandler struct {
	productRepository contracts.ProductRepository
}

func NewCreateProductCommandHandler(productRepository contracts.ProductRepository) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{productRepository: productRepository}
}

func (c *CreateProductCommandHandler) Handle(
	ctx context.Context,
	command *CreateProductCommand,
) (*dtos.CreateProductCommandResponse, error) {
	if command == nil {
		return nil, customErrors.NewApplicationErrorWithCode("command cannot be nil", http.StatusBadRequest)
	}

	if command.ProductID == uuid.Nil {
		return nil, customErrors.NewApplicationErrorWithCode("product id cannot be empty", http.StatusBadRequest)
	}

	product := MapCreateProductToProduct(command)

	createdProduct, err := c.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, customErrors.NewApplicationErrorWrap(
			err,
			"failed to create product in the repository",
		)
	}
	response := &dtos.CreateProductCommandResponse{ProductID: createdProduct.ProductID}

	// Publish notification event to the mediatr for dispatching to the notification handlers
	productCreatedEvent := mappings.MapProductToCreatedProduct(createdProduct)
	err = mediatr.Publish[*events.ProductCreatedEvent](ctx, productCreatedEvent)
	if err != nil {
		return nil, customErrors.NewApplicationErrorWrap(
			err,
			"error in publishing ProductCreatedEvent event",
		)
	}

	return response, nil
}
