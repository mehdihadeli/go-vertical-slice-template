package application

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/models"
)

func (a *Application) MigrateDatabase() error {
	return a.ResolveDependencyFunc(func(g *gorm.DB) error {
		// Auto-migrate the Person model
		err := g.AutoMigrate(&models.Product{})
		if err != nil {
			return err
		}

		g.Create(&models.Product{Name: "Test", CreatedAt: time.Now(), ProductID: uuid.NewV4(), Price: 100, Description: "Test description"})

		return nil
	})
}
