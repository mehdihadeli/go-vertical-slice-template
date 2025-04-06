package application

import (
	"time"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/models"

	"emperror.dev/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (a *Application) ConfigCatalogs() error {
	err := a.mapEndpoints()
	if err != nil {
		return errors.WrapIf(err, "Error in mapping endpoints")
	}

	err = a.migrateDatabase()
	if err != nil {
		return errors.WrapIf(err, "Error in migrating database")
	}

	return err
}

func (a *Application) mapEndpoints() error {
	a.ResolveRequiredDependencyFunc(func(endpoints []contracts.Endpoint) {
		for _, endpoint := range endpoints {
			endpoint.MapEndpoint()
		}
	})

	return nil
}

func (a *Application) migrateDatabase() error {
	return a.ResolveDependencyFunc(func(g *gorm.DB) error {
		// Auto-migrate the Person model
		err := g.AutoMigrate(&models.Product{})
		if err != nil {
			return err
		}

		g.Create(
			&models.Product{
				Name:        "Test",
				CreatedAt:   time.Now(),
				ProductID:   uuid.NewV4(),
				Price:       100,
				Description: "Test description",
			},
		)

		return nil
	})
}
