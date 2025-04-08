package config

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/dependency"
)

func AddEnv(dependencies *dependency.ServiceCollection, environments ...environemnt.Environment) {
	dependency.Add[environemnt.Environment](
		dependencies,
		func(sp *dependency.ServiceProvider) (environemnt.Environment, error) {
			return environemnt.ConfigEnv(environments...), nil
		},
	)
}
