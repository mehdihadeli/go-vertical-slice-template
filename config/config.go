package config

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
)

type Config struct {
	AppOptions      AppOptions      `mapstructure:"appOptions"      env:"AppOptions"`
	EchoHttpOptions EchoHttpOptions `mapstructure:"echoHttpOptions"`
}

func NewAppConfig(env environemnt.Environment) (*Config, error) {
	cfg, err := config.BindConfig[*Config](env)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

type AppOptions struct {
	Name string `mapstructure:"name" env:"Name"`
}

func (cfg *AppOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.Name)
}

func (cfg *AppOptions) GetMicroserviceName() string {
	return cfg.Name
}

type EchoHttpOptions struct {
	Port                string   `mapstructure:"port"                validate:"required" env:"Port"`
	Development         bool     `mapstructure:"development"                             env:"Development"`
	BasePath            string   `mapstructure:"basePath"            validate:"required" env:"BasePath"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"                     env:"DebugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
	Timeout             int      `mapstructure:"timeout"                                 env:"Timeout"`
	Host                string   `mapstructure:"host"                                    env:"Host"`
	Name                string   `mapstructure:"name"                                    env:"Name"`
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
