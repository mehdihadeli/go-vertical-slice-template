package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	applicationbuilder "github.com/go-vertical-slice-template/internal/app/application_builder"
	"github.com/go-vertical-slice-template/internal/shared/dependencies"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	builder := applicationbuilder.NewApplicationBuilder()

	err := builder.Services.Add(dependencies.Dependencies...)
	if err != nil {
		log.Fatal(err.Error())
	}

	app := builder.Build()

	// configure services
	err = app.ConfigMediator()
	if err != nil {
		log.Fatal("Error in setting mediator handlers", err)
	}

	app.MapEndpoints()

	app.Run(ctx)
}
