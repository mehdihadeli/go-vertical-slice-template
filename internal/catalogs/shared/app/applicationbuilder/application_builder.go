package applicationbuilder

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"
	"go.uber.org/dig"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app/application"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/constants"
	"github.com/spf13/viper"
)

type ApplicationBuilder struct {
	Container   *dig.Container
	Logger      logger.Logger
	Environment environemnt.Environment
	overrides   []*Override
}

type Override struct {
	DecoratorFunc interface{}
	Opts          []dig.DecorateOption
}

func NewApplicationBuilder(environments ...environemnt.Environment) *ApplicationBuilder {
	// Create the app container.
	// Do not forget to delete it at the end.
	// Create a Container with the default scopes (App, Request, SubRequest).
	container := dig.New()

	err := logger.AddLogger(container)
	if err != nil {
		log.Fatalln(err)
	}

	setConfigPath()
	err = config.AddEnv(container, environments...)
	if err != nil {
		log.Fatalln(err)
	}

	var l logger.Logger
	var env environemnt.Environment

	err = container.Invoke(func(logger logger.Logger, environment environemnt.Environment) error {
		env = environment
		l = logger

		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	appBuilder := &ApplicationBuilder{Container: container, Logger: l, Environment: env}

	return appBuilder
}

func (b *ApplicationBuilder) Build() *application.Application {
	// Apply overrides first
	for _, override := range b.overrides {
		err := b.Container.Decorate(override.DecoratorFunc, override.Opts...)
		if err != nil {
			b.Logger.Fatal(err)
		}
	}

	container := b.Container
	var app = application.NewApplication(container)

	return app
}

// WithOverride Can override test configs here, or use our seperated `TestApplicationBuilder`
func (b *ApplicationBuilder) WithOverride(decoratorFunc interface{}, opts ...dig.DecorateOption) *ApplicationBuilder {
	b.overrides = append(b.overrides, &Override{decoratorFunc, opts})
	return b
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
