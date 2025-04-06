package commands

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/models"
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreateProductCommand struct {
	ProductID   uuid.UUID `validate:"required"`
	Name        string    `validate:"required,gte=0,lte=255"`
	Description string    `validate:"required,gte=0,lte=5000"`
	Price       float64   `validate:"required,gte=0"`
	CreatedAt   time.Time `validate:"required"`
}

func NewCreateProductCommand(name string, description string, price float64) *CreateProductCommand {
	return &CreateProductCommand{ProductID: uuid.NewV4(), Name: name, Description: description, Price: price, CreatedAt: time.Now()}
}

func MapCreateProductToProduct(command *CreateProductCommand) *models.Product {
	product := &models.Product{
		ProductID:   command.ProductID,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		CreatedAt:   command.CreatedAt,
	}

	return product
}
