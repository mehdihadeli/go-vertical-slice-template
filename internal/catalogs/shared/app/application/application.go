package application

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
	"go.uber.org/zap"

	"github.com/go-vertical-slice-template/config"
)

type Application struct {
	Container di.Container
	Echo      *echo.Echo
	Logger    *zap.SugaredLogger
	Cfg       *config.Config
}

func NewApplication(container di.Container, echo *echo.Echo, logger *zap.SugaredLogger, cfg *config.Config) *Application {
	return &Application{Container: container, Echo: echo, Logger: logger, Cfg: cfg}
}

func (a *Application) Run() {
	//https://medium.com/@mokiat/proper-http-shutdown-in-go-bd3bfaade0f2
	defaultDuration := time.Second * 10

	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer ctxCancel()

	a.Start(ctx)

	<-ctx.Done()

	// The context is used to inform the server it has 10 seconds to finish
	// All Graceful shutdowns, should be in the `Stop` method  
	stopCtx, stopCtxCancel := context.WithTimeout(context.Background(), defaultDuration)
	defer stopCtxCancel()
	a.Stop(stopCtx)
}

func (a *Application) Start(ctx context.Context) {
	echoStartHook(ctx, a)
}

func (a *Application) Stop(shutdownCtx context.Context) {
	echoStopHook(shutdownCtx, a)

	a.Container.Delete()
	log.Println("Graceful shutdown complete.")
}

// Hooks
func echoStopHook(stopCtx context.Context, application *Application) {
	if err := application.Echo.Shutdown(stopCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
}

func echoStartHook(ctx context.Context, application *Application) {
	go func() {
		// When Shutdown is called, Serve, ListenAndServe, and ListenAndServeTLS immediately return ErrServerClosed. Make sure the program doesn't exit and waits instead for Shutdown to return.
		if err := application.Echo.Start(application.Cfg.EchoHttpOptions.Port); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()
}
