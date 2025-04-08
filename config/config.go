package config

import (
	"strings"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	typeMapper "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/reflection/typemapper"

	"github.com/iancoleman/strcase"
)

type AppOptions struct {
	Name string `mapstructure:"name" env:"Name"`
}

func NewAppOptions(environment environemnt.Environment) (*AppOptions, error) {
	optionName := strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[AppOptions]())
	cfg, err := config.BindConfigKey[*AppOptions](optionName, environment)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *AppOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.Name)
}

func (cfg *AppOptions) GetMicroserviceName() string {
	return cfg.Name
}
