package logger

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"go.uber.org/dig"
)

func AddLogger(container *dig.Container) error {
	err := container.Provide(func(environment environemnt.Environment) (*LogOptions, error) {
		return ProvideLogConfig(environment)
	})
	if err != nil {
		return err
	}

	err = container.Provide(func(opts *LogOptions, environment environemnt.Environment) Logger {
		return NewZapLogger(opts, environment)
	})
	if err != nil {
		return err
	}
	
	return err
}
