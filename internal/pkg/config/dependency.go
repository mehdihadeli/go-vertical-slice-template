package config

import (
	"go.uber.org/dig"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
)

func AddEnv(container *dig.Container, environments ...environemnt.Environment) error {
	err := container.Provide(func() environemnt.Environment {
		return environemnt.ConfigEnv(environments...)
	})

	return err
}
