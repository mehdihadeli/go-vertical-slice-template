package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	di_cqrs "github.com/go-vertical-slice-template"
	applicationbuilder "github.com/go-vertical-slice-template/internal/app/application_builder"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	builder := applicationbuilder.NewApplicationBuilder()

	err := builder.Services.Add(di_cqrs.Dependencies...)
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
