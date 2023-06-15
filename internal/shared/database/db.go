package database

import (
	"time"

	"github.com/glebarez/sqlite"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"github.com/go-vertical-slice-template/internal/products/models"
)

func NewGormDB() (*gorm.DB, error) {
	// https://gorm.io/docs/connecting_to_the_database.html#SQLite
	// https://github.com/glebarez/sqlite
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MigrateDb(db *gorm.DB) error {
	// Auto-migrate the Person model
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		return err
	}

	db.Create(&models.Product{Name: "Test", CreatedAt: time.Now(), ProductID: uuid.NewV4(), Price: 100, Description: "Test description"})

	return nil
}
