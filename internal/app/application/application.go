package application

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
)

type Application struct {
	Container di.Container
	Echo      *echo.Echo
	Logger    *zap.SugaredLogger
}

func (a *Application) Run(ctx context.Context) {
	defer a.Container.Delete()

	err := a.Echo.Start(":9080")

	if err != nil {
		a.Logger.Fatal("Error starting Server", err)
	}

	<-ctx.Done()

	if err := a.Echo.Shutdown(ctx); err != nil {
		a.Logger.Fatal("(Shutdown) err", err)
	}
}
