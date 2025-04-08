package echoweb

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/dependency"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/echoweb/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"

	"github.com/labstack/echo/v4"
)

func AddEcho(dependencies *dependency.ServiceCollection) {
	dependency.Add[*config.EchoHttpOptions](
		dependencies,
		func(sp *dependency.ServiceProvider) (*config.EchoHttpOptions, error) {
			environment := dependency.GetGenericRequiredService[environemnt.Environment](sp)
			return config.ConfigEchoOptions(environment)
		},
	)

	dependency.Add[*echo.Echo](dependencies, func(sp *dependency.ServiceProvider) (*echo.Echo, error) {
		l := dependency.GetGenericRequiredService[logger.Logger](sp)
		return NewEcho(l)
	})
}
