package app

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app/application"
	applicationbuilder "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app/applicationbuilder"
)

type App struct{}

// NewApp creates a new App
func NewApp() *App {
	app := &App{}
	return app
}

func (a *App) Run() {
	builder := createApplicationBuilder()

	app := builder.Build()

	configureApplication(app)

	app.Run()
}

func configureApplication(app *application.Application) {
	// configure services
	err := app.ConfigInfrastructure()
	if err != nil {
		app.Logger.Fatal(err)
	}
	err = app.ConfigCatalogs()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func createApplicationBuilder() *applicationbuilder.ApplicationBuilder {
	builder := applicationbuilder.NewApplicationBuilder()

	builder.AddInfrastructure()
	err := builder.AddCatalogs()
	if err != nil {
		builder.Logger.Fatal(err)
	}

	return builder
}
