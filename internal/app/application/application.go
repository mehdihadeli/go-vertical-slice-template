package application

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type Application struct {
	Container di.Container
	Echo      *echo.Echo
}

func (a *Application) Run(ctx context.Context) {
	defer a.Container.Delete()

	err := a.Echo.Start(":9080")

	if err != nil {
		log.Fatal("Error starting Server", err)
	}

	<-ctx.Done()

	if err := a.Echo.Shutdown(ctx); err != nil {
		log.Fatal("(Shutdown) err", err)
	}
}
