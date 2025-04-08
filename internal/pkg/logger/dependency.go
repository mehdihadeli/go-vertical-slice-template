package logger

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/dependency"
)

func AddLogger(dependencies *dependency.ServiceCollection, environment environemnt.Environment) {
	dependency.Add[*LogOptions](dependencies, func(sp *dependency.ServiceProvider) (*LogOptions, error) {
		return ConfigLopOptions(environment)
	})

	dependency.Add[Logger](dependencies, func(sp *dependency.ServiceProvider) (Logger, error) {
		options := dependency.GetGenericRequiredService[*LogOptions](sp)
		return NewZapLogger(environment, options)
	})
}
