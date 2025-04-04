package options

import (
	"fmt"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config"
)

type GormOptions struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"dbName"`
	SSLMode  bool   `mapstructure:"sslMode"`
	Password string `mapstructure:"password"`
}

func (h *GormOptions) Dns() string {
	datasource := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		h.User,
		h.Password,
		h.Host,
		h.Port,
		"postgres",
	)

	return datasource
}

func ProvideConfig() (*GormOptions, error) {
	return config.BindConfigKey[*GormOptions]("gormOptions")
}
