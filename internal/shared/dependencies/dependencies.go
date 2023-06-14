package dependencies

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"

	productApi "github.com/go-vertical-slice-template/internal/products/api"
	"github.com/go-vertical-slice-template/internal/products/repository"
)

var Dependencies = []di.Def{
	{
		Name:  "echo",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return echo.New(), nil
		},
	},
	{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return log.Logger{}, nil
		},
	},
	{
		Name:  "productController",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			echo := ctn.Get("echo").(*echo.Echo)
			return productApi.NewProductsController(echo), nil
		},
	},
	{
		Name:  "productRepository",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return repository.NewInMemoryProductRepository(), nil
		},
	},
}
