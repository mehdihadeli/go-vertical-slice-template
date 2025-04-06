package mappings

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/events"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/models"
)

func MapProductToProductDto(product *models.Product) *dtos.ProductDto {
	return &dtos.ProductDto{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		UpdatedAt:   product.UpdatedAt,
		CreatedAt:   product.CreatedAt,
		Price:       product.Price,
	}
}

func MapProductToCreatedProduct(createdProduct *models.Product) *events.ProductCreatedEvent {
	productCreatedEvent := events.NewProductCreatedEvent(createdProduct.ProductID, createdProduct.Name, createdProduct.Description, createdProduct.Price, createdProduct.CreatedAt)

	return productCreatedEvent
}
