package echoweb

import (
	"strings"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/constants"
	handlers "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/echoweb/hadnlers"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/echoweb/middlewares/log"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/echoweb/middlewares/problemdetail"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho(logger logger.Logger) (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true

	skipper := func(c echo.Context) bool {
		return strings.Contains(c.Request().URL.Path, "swagger") ||
			strings.Contains(c.Request().URL.Path, "metrics") ||
			strings.Contains(c.Request().URL.Path, "health") ||
			strings.Contains(c.Request().URL.Path, "favicon.ico")
	}

	// set error handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// bypass skip endpoints and its error
		if skipper(c) {
			return
		}

		handlers.ProblemDetailErrorHandlerFunc(err, c, logger)
	}

	// log errors and information
	e.Use(
		log.EchoLogger(
			logger,
			log.WithSkipper(skipper),
		),
	)
	e.Use(middleware.BodyLimit(constants.BodyLimit))
	e.Use(middleware.RequestID())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level:   constants.GzipLevel,
		Skipper: skipper,
	}))
	// should be last middleware
	e.Use(problemdetail.ProblemDetail(problemdetail.WithSkipper(skipper)))

	return e, nil
}
