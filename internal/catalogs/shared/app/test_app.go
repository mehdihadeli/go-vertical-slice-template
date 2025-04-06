package app

import (
	"testing"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app/application"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app/applicationbuilder"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/constants"

	"github.com/spf13/viper"
	"go.uber.org/dig"
)

type (
	BuilderOption       func(*applicationbuilder.ApplicationBuilder)
	ConfigurationOption func(*application.Application)
)

type TestApp struct {
	*App
	overrides []*applicationbuilder.Override
}

// NewTestApp creates a new App for test with optional configurations
func NewTestApp() *TestApp {
	viper.Set(constants.AppEnv, environemnt.Test.GetEnvironmentName())
	app := &TestApp{App: &App{}}

	return app
}

func (a *TestApp) RunTest(t *testing.T) *dig.Container {
	builder := createApplicationBuilder()
	// Apply override builder options
	for _, opt := range a.overrides {
		builder.WithOverride(opt.DecoratorFunc, opt.Opts...)
	}

	app := builder.Build()

	configureApplication(app)

	container := app.Container

	app.RunTest(t)

	return container
}

// WithOverrideBuilder Can override test configs here, or use our seperated `TestApplicationBuilder`
func (a *TestApp) WithOverrideBuilder(decoratorFunc interface{}, opts ...dig.DecorateOption) *TestApp {
	a.overrides = append(a.overrides, &applicationbuilder.Override{DecoratorFunc: decoratorFunc, Opts: opts})

	return a
}
