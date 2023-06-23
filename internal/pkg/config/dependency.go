package config

import (
	"github.com/sarulabs/di"

	"github.com/go-vertical-slice-template/internal/pkg/config/environemnt"
)

func AddEnv(container *di.Builder) error {
	envDep := di.Def{
		Name:  "env",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return environemnt.ConfigAppEnv(), nil
		},
	}
	return container.Add(envDep)
}
