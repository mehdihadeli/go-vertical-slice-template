package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/database/options"
)

func NewGormDB(options *options.GormOptions) (*gorm.DB, error) {
	// https://gorm.io/docs/connecting_to_the_database.html#SQLite
	// https://github.com/glebarez/sqlite
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
