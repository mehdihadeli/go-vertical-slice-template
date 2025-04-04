package main

import applicationbuilder "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app/application_builder"

func main() {
	builder := applicationbuilder.NewApplicationBuilder()

	builder.AddCore()
	builder.AddInfrastructure()
	builder.AddRepositories()
	builder.AddRoutes()

	app := builder.Build()

	// configure services
	err := app.MigrateDatabase()
	if err != nil {
		app.Logger.Fatal("Error in migrating database", err)
	}

	err = app.ConfigMediator()
	if err != nil {
		app.Logger.Fatal("Error in setting mediator handlers", err)
	}

	app.MapEndpoints()

	app.ConfigSwagger()

	app.Run()
}
