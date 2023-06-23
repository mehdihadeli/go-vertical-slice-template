package application

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"github.com/go-vertical-slice-template/internal/catalogs/products/models"
)

func (a *Application) MigrateDatabase() error {
	db := a.Container.Get("gorm").(*gorm.DB)
	// Auto-migrate the Person model
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		return err
	}

	db.Create(&models.Product{Name: "Test", CreatedAt: time.Now(), ProductID: uuid.NewV4(), Price: 100, Description: "Test description"})

	return nil
}
