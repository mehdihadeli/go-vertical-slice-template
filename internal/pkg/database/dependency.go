package database

import (
	"go.uber.org/dig"
	"gorm.io/gorm"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/database/options"
)

func AddGorm(container *dig.Container) error {
	err := container.Provide(func() (*options.GormOptions, error) {
		return options.ProvideConfig()
	})

	err = container.Provide(func(opts *options.GormOptions) (*gorm.DB, error) {
		return NewGorm(opts)
	})

	return err
}
