package options

import (
	"fmt"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
)

type GormOptions struct {
	Host        string `mapstructure:"host"        env:"GORM_OPTIONS_HOST"`
	Port        int    `mapstructure:"port"        env:"GORM_OPTIONS_PORT"`
	User        string `mapstructure:"user"`
	DBName      string `mapstructure:"dbName"`
	SSLMode     bool   `mapstructure:"sslMode"`
	Password    string `mapstructure:"password"`
	UseInMemory bool   `mapstructure:"useInMemory"`
	UseSQLLite  bool   `mapstructure:"useSqlLite"`
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

func ConfigGormOptions(environment environemnt.Environment) (*GormOptions, error) {
	return config.BindConfigKey[*GormOptions]("gormOptions", environment)
}
