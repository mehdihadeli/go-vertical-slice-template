package applicationbuilder

import (
	echo2 "github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"go.uber.org/zap"

	"github.com/go-vertical-slice-template/internal/app/application"
)

type ApplicationBuilder struct {
	Services *di.Builder
	Logger   *zap.SugaredLogger
}

func NewApplicationBuilder() *ApplicationBuilder {
	log := createLogger()

	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ApplicationBuilder{Services: builder, Logger: log}
}

func (b *ApplicationBuilder) Build() *application.Application {
	container := b.Services.Build()

	echo := container.Get("echo").(*echo2.Echo)
	return &application.Application{Container: container, Echo: echo, Logger: b.Logger}
}

func createLogger() *zap.SugaredLogger {
	// https://github.com/uber-go/zap#quick-start
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	log := logger.Sugar()

	return log
}
