package application

import (
	"gorm.io/gorm"

	"github.com/go-vertical-slice-template/internal/shared/database"
)

func (a *Application) MigrateDatabase() error {
	g := a.Container.Get("gorm").(*gorm.DB)
	return database.MigrateDb(g)
}
