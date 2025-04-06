package config

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"

	"go.uber.org/dig"
)

func AddAppConfig(container *dig.Container) error {
	err := container.Provide(func(environment environemnt.Environment) (*Config, error) {
		return NewAppConfig(environment)
	})

	return err
}
