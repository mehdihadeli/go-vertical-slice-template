package application

import (
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/mehdihadeli/go-vertical-slice-template/docs"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
)

func (a *Application) MapEndpoints() {
	endpoints := a.Container.Get("routes").([]contracts.Endpoint)
	for _, endpoint := range endpoints {
		endpoint.MapEndpoint()
	}

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Catalogs Write-Service Api"
	docs.SwaggerInfo.Description = "Catalogs Write-Service Api."

	a.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
