package applicationbuilder

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mehdihadeli/go-vertical-slice-template/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app/application"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/constants"
	config2 "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/echoweb/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Override func(builder *ApplicationBuilder) error

type ApplicationBuilder struct {
	Logger            logger.Logger
	Environment       environemnt.Environment
	appOptions        *config.AppOptions
	echoOptions       *config2.EchoHttpOptions
	gormDB            *gorm.DB
	echo              *echo.Echo
	endpoints         []contracts.Endpoint
	productRepository contracts.ProductRepository
}

func NewApplicationBuilder(environments ...environemnt.Environment) *ApplicationBuilder {
	env := environemnt.ConfigEnv(environments...)
	setConfigPath()

	lopOptions, err := logger.ConfigLopOptions(env)
	if err != nil {
		log.Fatal(err)
	}
	l, err := logger.NewZapLogger(env, lopOptions)
	if err != nil {
		log.Fatal(err)
	}

	appBuilder := &ApplicationBuilder{Logger: l, Environment: env}

	return appBuilder
}

func (b *ApplicationBuilder) Build() *application.Application {
	app := application.NewApplication(
		b.Environment,
		b.Logger,
		b.gormDB,
		b.echo,
		b.echoOptions,
		b.productRepository,
		b.endpoints,
	)

	return app
}

func setConfigPath() {
	// https://stackoverflow.com/a/47785436/581476
	wd, _ := os.Getwd()

	// https://articles.wesionary.team/environment-variable-configuration-in-your-golang-project-using-viper-4e8289ef664d
	// get from `os env` or viper `internal registry`
	pn := viper.Get(constants.PROJECT_NAME_ENV)
	if pn == nil {
		return
	}
	for !strings.HasSuffix(wd, pn.(string)) {
		wd = filepath.Dir(wd)
	}

	// Get the absolute path of the executed project directory
	absCurrentDir, _ := filepath.Abs(wd)

	viper.Set(constants.AppRootPath, absCurrentDir)

	// Get the path to the "config" folder within the project directory
	configPath := filepath.Join(absCurrentDir, "config")

	viper.Set(constants.ConfigPath, configPath)
}
