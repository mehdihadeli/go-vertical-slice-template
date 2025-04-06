package contracts

import (
	"context"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/models"

	uuid "github.com/satori/go.uuid"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	GetProductById(ctx context.Context, uuid uuid.UUID) (*models.Product, error)
}
