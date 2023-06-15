package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	applicationbuilder "github.com/go-vertical-slice-template/internal/app/application_builder"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	builder := applicationbuilder.NewApplicationBuilder()

	builder.AddLogger()
	builder.AddEcho()
	builder.AddGorm()
	builder.AddRepositories()
	builder.AddControllers()

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

	app.Run(ctx)
}
