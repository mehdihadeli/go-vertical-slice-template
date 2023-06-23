package commands

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"

	"github.com/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/go-vertical-slice-template/internal/catalogs/products/features/creating_product/dtos"
	"github.com/go-vertical-slice-template/internal/catalogs/products/features/creating_product/events"
	"github.com/go-vertical-slice-template/internal/catalogs/products/models"
)

type CreateProductCommandHandler struct {
	productRepository contracts.ProductRepository
}

func NewCreateProductCommandHandler(productRepository contracts.ProductRepository) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{productRepository: productRepository}
}

func (c *CreateProductCommandHandler) Handle(ctx context.Context, command *CreateProductCommand) (*dtos.CreateProductCommandResponse, error) {

	product := &models.Product{
		ProductID:   command.ProductID,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		CreatedAt:   command.CreatedAt,
	}

	createdProduct, err := c.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	response := &dtos.CreateProductCommandResponse{ProductID: createdProduct.ProductID}

	// Publish notification event to the mediatr for dispatching to the notification handlers

	productCreatedEvent := events.NewProductCreatedEvent(createdProduct.ProductID, createdProduct.Name, createdProduct.Description, createdProduct.Price, createdProduct.CreatedAt)
	err = mediatr.Publish[*events.ProductCreatedEvent](ctx, productCreatedEvent)
	if err != nil {
		return nil, err
	}

	return response, nil
}
