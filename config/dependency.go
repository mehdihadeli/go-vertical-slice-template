package config

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/dependency"
)

func AddAppConfig(dependencies *dependency.ServiceCollection, environment environemnt.Environment) {
	dependency.Add[*AppOptions](dependencies, func(sp *dependency.ServiceProvider) (*AppOptions, error) {
		return NewAppOptions(environment)
	})
}
