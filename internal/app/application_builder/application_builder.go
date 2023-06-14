package applicationbuilder

import (
	"log"

	echo2 "github.com/labstack/echo/v4"
	"github.com/sarulabs/di"

	"github.com/go-vertical-slice-template/internal/app/application"
)

type ApplicationBuilder struct {
	Services *di.Builder
}

func NewApplicationBuilder() *ApplicationBuilder {
	// Create the app container.
	// Do not forget to delete it at the end.
	// Create a Services with the default scopes (App, Request, SubRequest).
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ApplicationBuilder{Services: builder}
}

func (b *ApplicationBuilder) Build() *application.Application {
	container := b.Services.Build()

	echo := container.Get("echo").(*echo2.Echo)
	return &application.Application{Container: container, Echo: echo}
}
