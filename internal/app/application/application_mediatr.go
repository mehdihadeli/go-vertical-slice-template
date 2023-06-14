package application

import (
	"github.com/mehdihadeli/go-mediatr"

	"github.com/go-vertical-slice-template/internal/products/contracts"
	"github.com/go-vertical-slice-template/internal/products/features/creating_product/commands"
	creatingProductsDtos "github.com/go-vertical-slice-template/internal/products/features/creating_product/dtos"
	"github.com/go-vertical-slice-template/internal/products/features/creating_product/events"
	gettingProductByIdDtos "github.com/go-vertical-slice-template/internal/products/features/getting_product_by_id/dtos"
	"github.com/go-vertical-slice-template/internal/products/features/getting_product_by_id/queries"
	"github.com/go-vertical-slice-template/internal/shared/behaviours"
)

func (a *Application) ConfigMediator() error {
	productRepository := a.Container.Get("productRepository").(contracts.ProductRepository)

	loggerPipeline := &behaviours.RequestLoggerBehaviour{}
	err := mediatr.RegisterRequestPipelineBehaviors(loggerPipeline)

	if err != nil {
		return err
	}

	createProductCommandHandler := commands.NewCreateProductCommandHandler(productRepository)
	err = mediatr.RegisterRequestHandler[*commands.CreateProductCommand, *creatingProductsDtos.CreateProductCommandResponse](createProductCommandHandler)
	if err != nil {
		return err
	}

	getByIdQueryHandler := queries.NewGetProductByIdHandler(productRepository)
	err = mediatr.RegisterRequestHandler[*queries.GetProductByIdQuery, *gettingProductByIdDtos.GetProductByIdQueryResponse](getByIdQueryHandler)
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
