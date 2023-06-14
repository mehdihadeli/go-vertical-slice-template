package application

import (
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/go-vertical-slice-template/docs"
	productApi "github.com/go-vertical-slice-template/internal/products/api"
)

func (a *Application) MapEndpoints() {
	productController := a.Container.Get("productController").(*productApi.ProductsController)

	productApi.MapProductsRoutes(productController)

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Catalogs Write-Service Api"
	docs.SwaggerInfo.Description = "Catalogs Write-Service Api."

	a.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
