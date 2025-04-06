package application

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"
	"go.uber.org/dig"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/mehdihadeli/go-vertical-slice-template/config"
	"github.com/pkg/errors"
)

type Application struct {
	Container *dig.Container
	Echo      *echo.Echo
	Logger    logger.Logger
	Cfg       *config.Config
}

func NewApplication(container *dig.Container) *Application {
	app := &Application{}
	err := container.Invoke(func(c *config.Config, e *echo.Echo, logger logger.Logger) error {
		app.Container = container
		app.Echo = e
		app.Logger = logger
		app.Cfg = c

		return nil
	})

	if err != nil {
		app.Logger.Fatal(err)
	}

	return app
}

func (a *Application) ResolveDependencyFunc(function interface{}) error {
	return a.Container.Invoke(function)
}

func (a *Application) ResolveRequiredDependencyFunc(function interface{}) {
	err := a.Container.Invoke(function)
	if err != nil {
		panic(fmt.Sprintf("failed to resolve dependency: %v", err))
	}
}

func (a *Application) Run() {
	//https://dev.to/mokiat/proper-http-shutdown-in-go-3fji
	// https://github.com/uber-go/fx/blob/master/app_test.go
	defaultDuration := time.Second * 20

	// short context timeout just for starting `Start hooks` and setup dependencies
	startCtx, cancel := context.WithTimeout(context.Background(), defaultDuration)
	defer cancel()
	a.Start(startCtx)

	// block the main goroutine and keep the app running until an interrupt signal (SIGINT / SIGTERM) is received.
	<-a.Wait()

	// short context timeout just for doing `Stop hooks` and a graceful shutdown
	// The context is used to inform the server it has 10 seconds to finish
	// All Graceful shutdowns, should be in the `Stop` method
	stopCtx, stopCancellation := context.WithTimeout(context.Background(), defaultDuration)
	defer stopCancellation()
	a.Stop(stopCtx)
}

func (a *Application) RunTest(t *testing.T) {
	// we need a longer timout for up and running our testcontainers
	duration := time.Second * 300

	// short timeout for handling start hooks and setup dependencies
	startCtx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	a.Start(startCtx)

	// we don't need wait, and we don't want to block goroutine here and test should continue, and we stop server in end of test

	t.Cleanup(func() {
		// short context timeout just for doing `Stop hooks` and a graceful shutdown
		// The context is used to inform the server it has 10 seconds to finish
		// All Graceful shutdowns, should be in the `Stop` method
		stopCtx, cancel := context.WithTimeout(context.Background(), duration)
		defer cancel()

		a.Stop(stopCtx)
	})
}

func (a *Application) Start(startCtx context.Context) {
	// start hooks
	echoStartHook(startCtx, a)
}

func (a *Application) Stop(shutdownCtx context.Context) {
	// stop hooks
	echoStopHook(shutdownCtx, a)

	log.Println("Graceful shutdown complete.")
}

func (a *Application) Wait() <-chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	return sigChan
}

// Hooks
func echoStopHook(stopCtx context.Context, application *Application) {
	if err := application.Echo.Shutdown(stopCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
}

func echoStartHook(startCtx context.Context, application *Application) {
	go func() {
		// When Shutdown is called, Serve, ListenAndServe, and ListenAndServeTLS immediately return ErrServerClosed. Make sure the program doesn't exit and waits instead for Shutdown to return.
		if err := application.Echo.Start(application.Cfg.EchoHttpOptions.Port); !errors.Is(err, http.ErrServerClosed) {
			application.Logger.Fatalf("HTTP server error: %v", err)
		}
		application.Logger.Info("Stopped serving new HTTP connections.")
	}()
}
