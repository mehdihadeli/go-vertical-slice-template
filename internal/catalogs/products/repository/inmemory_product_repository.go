package repository

import (
	"context"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"

	"gorm.io/gorm"

	uuid "github.com/satori/go.uuid"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/models"
)

type InMemoryProductRepository struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewInMemoryProductRepository(db *gorm.DB, logger logger.Logger) contracts.ProductRepository {
	return &InMemoryProductRepository{db: db, logger: logger}
}

func (p *InMemoryProductRepository) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	err := p.db.WithContext(ctx).Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *InMemoryProductRepository) GetProductById(ctx context.Context, uuid uuid.UUID) (*models.Product, error) {
	var product *models.Product
	err := p.db.WithContext(ctx).First(&product, "product_id = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
