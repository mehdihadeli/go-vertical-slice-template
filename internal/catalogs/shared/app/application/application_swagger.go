package application

import (
	"github.com/mehdihadeli/go-vertical-slice-template/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (a *Application) ConfigSwagger() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Catalogs Write-Service Api"
	docs.SwaggerInfo.Description = "Catalogs Write-Service Api."

	a.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
