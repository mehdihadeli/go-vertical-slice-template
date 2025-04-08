package application

import (
	"github.com/mehdihadeli/go-vertical-slice-template/docs"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/commands"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/events"
	dtos2 "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/queries"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/behaviours"

	"github.com/cockroachdb/errors"
	"github.com/mehdihadeli/go-mediatr"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (a *Application) ConfigInfrastructure() error {
	err := a.configMediator()
	if err != nil {
		return errors.Wrap(err, "Error in setting mediator handlers")
	}

	a.configSwagger()

	return err
}

func (a *Application) configMediator() error {
	loggerPipeline := &behaviours.RequestLoggerBehaviour{}
	err := mediatr.RegisterRequestPipelineBehaviors(loggerPipeline)
	if err != nil {
		return err
	}

	createProductCommandHandler := commands.NewCreateProductCommandHandler(a.ProductRepository)
	err = mediatr.RegisterRequestHandler[*commands.CreateProductCommand, *dtos.CreateProductCommandResponse](
		createProductCommandHandler,
	)
	if err != nil {
		return err
	}

	getByIdQueryHandler := queries.NewGetProductByIdHandler(a.ProductRepository)
	err = mediatr.RegisterRequestHandler[*queries.GetProductByIdQuery, *dtos2.GetProductByIdQueryResponse](
		getByIdQueryHandler,
	)
	if err != nil {
		return err
	}

	notificationHandler := events.NewProductCreatedEventHandler()
	err = mediatr.RegisterNotificationHandler[*events.ProductCreatedEvent](notificationHandler)
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) configSwagger() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Catalogs Write-Service Api"
	docs.SwaggerInfo.Description = "Catalogs Write-Service Api."

	a.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
