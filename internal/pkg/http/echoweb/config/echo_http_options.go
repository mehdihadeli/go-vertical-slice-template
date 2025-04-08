package config

import (
	"fmt"
	"net/url"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	typeMapper "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/reflection/typemapper"

	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[EchoHttpOptions]())

type EchoHttpOptions struct {
	Port                string   `mapstructure:"port"                validate:"required" env:"ECHO_HTTP_OPTIONS_PORT"`
	Development         bool     `mapstructure:"development"                             env:"ECHO_HTTP_OPTIONS_DEVELOPMENT"`
	BasePath            string   `mapstructure:"basePath"            validate:"required" env:"ECHO_HTTP_OPTIONS_BASE_PATH"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
	Timeout             int      `mapstructure:"timeout"                                 env:"ECHO_HTTP_OPTIONS_Timeout"`
	Host                string   `mapstructure:"host"                                    env:"ECHO_HTTP_OPTIONS_Host"`
	Name                string   `mapstructure:"name"                                    env:"ECHO_HTTP_OPTIONS_Name"`
}

func (c *EchoHttpOptions) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}

func (c *EchoHttpOptions) BasePathAddress() string {
	path, err := url.JoinPath(c.Address(), c.BasePath)
	if err != nil {
		return ""
	}
	return path
}

func ConfigEchoOptions(environment environemnt.Environment) (*EchoHttpOptions, error) {
	return config.BindConfigKey[*EchoHttpOptions]("echoHttpOptions", environment)
}
