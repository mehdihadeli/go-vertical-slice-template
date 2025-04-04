package echoweb

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts/params"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func AddEcho(container *dig.Container) error {
	err := container.Provide(func() *echo.Echo {
		return echo.New()
	})
	if err != nil {
		return err
	}

	err = container.Provide(func(e *echo.Echo, l *zap.SugaredLogger) (*params.ProductRouteParams, error) {
		v1 := e.Group("/api/v1")
		products := v1.Group("/products")

		productsRouteParams := &params.ProductRouteParams{
			Logger:        l,
			Validator:     validator.New(),
			ProductsGroup: products,
		}

		return productsRouteParams, nil
	})

	if err != nil {
		return err
	}

	return nil
}
