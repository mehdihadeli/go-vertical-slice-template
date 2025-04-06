package config

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"

	"go.uber.org/dig"
)

func AddEnv(container *dig.Container, environments ...environemnt.Environment) error {
	err := container.Provide(func() environemnt.Environment {
		return environemnt.ConfigEnv(environments...)
	})

	return err
}
