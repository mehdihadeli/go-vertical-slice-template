package app

import (
	"testing"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app/application"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app/applicationbuilder"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/constants"

	"github.com/spf13/viper"
)

type (
	BuilderOption       func(*applicationbuilder.ApplicationBuilder)
	ConfigurationOption func(*application.Application)
)

type TestApp struct {
	*App
	overrides []applicationbuilder.Override
}

// NewTestApp creates a new App for test with optional configurations
func NewTestApp() *TestApp {
	viper.Set(constants.AppEnv, environemnt.Test.GetEnvironmentName())
	app := &TestApp{App: &App{}}

	return app
}

func (a *TestApp) RunTest(t *testing.T) *application.Application {
	builder := createApplicationBuilder()

	app := builder.Build()

	configureApplication(app)

	app.RunTest(t)

	return app
}
