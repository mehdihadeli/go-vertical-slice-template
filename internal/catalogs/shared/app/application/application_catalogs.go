package application

import (
	"time"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/models"

	"github.com/cockroachdb/errors"
	uuid "github.com/satori/go.uuid"
)

func (a *Application) ConfigCatalogs() error {
	err := a.mapEndpoints()
	if err != nil {
		return errors.Wrap(err, "Error in mapping endpoints")
	}

	err = a.migrateDatabase()
	if err != nil {
		return errors.Wrap(err, "Error in migrating database")
	}

	return err
}

func (a *Application) mapEndpoints() error {
	for _, endpoint := range a.Endpoints {
		endpoint.MapEndpoint()
	}

	return nil
}

func (a *Application) migrateDatabase() error {
	// Auto-migrate the Person model
	err := a.GormDB.AutoMigrate(&models.Product{})
	if err != nil {
		return err
	}

	a.GormDB.Create(
		&models.Product{
			Name:        "Test",
			CreatedAt:   time.Now(),
			ProductID:   uuid.NewV4(),
			Price:       100,
			Description: "Test description",
		},
	)

	return nil
}
