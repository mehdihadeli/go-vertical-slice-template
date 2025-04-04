package application

import (
	"github.com/mehdihadeli/go-mediatr"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creating_product/commands"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creating_product/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creating_product/events"
	dtos2 "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/getting_product_by_id/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/getting_product_by_id/queries"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/behaviours"
)

func (a *Application) ConfigMediator() error {
	productRepository := a.Container.Get("productRepository").(contracts.ProductRepository)

	loggerPipeline := &behaviours.RequestLoggerBehaviour{}
	err := mediatr.RegisterRequestPipelineBehaviors(loggerPipeline)

	if err != nil {
		return err
	}

	createProductCommandHandler := commands.NewCreateProductCommandHandler(productRepository)
	err = mediatr.RegisterRequestHandler[*commands.CreateProductCommand, *dtos.CreateProductCommandResponse](createProductCommandHandler)
	if err != nil {
		return err
	}

	getByIdQueryHandler := queries.NewGetProductByIdHandler(productRepository)
	err = mediatr.RegisterRequestHandler[*queries.GetProductByIdQuery, *dtos2.GetProductByIdQueryResponse](getByIdQueryHandler)
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
