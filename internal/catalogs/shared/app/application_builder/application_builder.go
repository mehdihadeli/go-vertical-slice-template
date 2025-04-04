package applicationbuilder

import (
	"go.uber.org/dig"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app/application"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/constants"
)

type ApplicationBuilder struct {
	Services    *dig.Container
	Logger      *zap.SugaredLogger
	Environment environemnt.Environment
}

func NewApplicationBuilder(environments ...environemnt.Environment) *ApplicationBuilder {
	env := environemnt.ConfigAppEnv(environments...)

	log := createLogger()
	setConfigPath()

	// Create the app container.
	// Do not forget to delete it at the end.
	// Create a Services with the default scopes (App, Request, SubRequest).
	builder := dig.New()
	return &ApplicationBuilder{Services: builder, Logger: log, Environment: env}
}

func (b *ApplicationBuilder) Build() *application.Application {
	container := b.Services
	var app = application.NewApplication(container)

	return app
}

func createLogger() *zap.SugaredLogger {
	// https://github.com/uber-go/zap#quick-start
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	log := logger.Sugar()

	return log
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
