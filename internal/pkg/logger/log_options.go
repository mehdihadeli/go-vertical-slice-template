package logger

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/reflection/typemapper"

	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typemapper.GetGenericTypeNameByT[LogOptions]())

type LogOptions struct {
	LogLevel      string `mapstructure:"level"`
	CallerEnabled bool   `mapstructure:"callerEnabled"`
}

func ConfigLopOptions(env environemnt.Environment) (*LogOptions, error) {
	return config.BindConfigKey[*LogOptions](optionName, env)
}
