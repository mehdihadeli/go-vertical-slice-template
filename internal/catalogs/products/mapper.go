package products

import (
	"github.com/go-vertical-slice-template/internal/catalogs/products/dtos"
	"github.com/go-vertical-slice-template/internal/catalogs/products/models"
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
