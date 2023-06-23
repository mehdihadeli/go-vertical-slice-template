package contracts

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"github.com/go-vertical-slice-template/internal/catalogs/products/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	GetProductById(ctx context.Context, uuid uuid.UUID) (*models.Product, error)
}
