package applicationbuilder

import (
	"emperror.dev/errors"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts/params"
	creatingproductendpoint "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/endpoints"
	gettingproductbyidendpoint "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/endpoints"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/repository"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"
	"gorm.io/gorm"
)

func (b *ApplicationBuilder) AddCatalogs() error {
	err := b.addRoutes()
	if err != nil {
		return err
	}

	err = b.addRepositories()
	if err != nil {
		return err
	}

	return nil
}

func (b *ApplicationBuilder) addRoutes() error {
	// https://echo.labstack.com/docs/routing
	err := b.Container.Provide(func(e *echo.Echo, l logger.Logger) (*params.ProductRouteParams, error) {
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
		return errors.WrapIf(err, "Error in mapping endpoints")
	}

	err = b.Container.Provide(func(productRouteParams *params.ProductRouteParams) ([]contracts.Endpoint, error) {
		createProductEndpoint := creatingproductendpoint.NewCreteProductEndpoint(productRouteParams)
		getProductById := gettingproductbyidendpoint.NewGetProductByIdEndpoint(productRouteParams)
		endpoints := []contracts.Endpoint{createProductEndpoint, getProductById}

		return endpoints, nil
	})

	if err != nil {
		return errors.WrapIf(err, "Error in mapping endpoints")
	}

	return nil
}

func (b *ApplicationBuilder) addRepositories() error {
	err := b.Container.Provide(func(g *gorm.DB, l logger.Logger) (contracts.ProductRepository, error) {
		return repository.NewInMemoryProductRepository(g, l), nil
	})

	if err != nil {
		return errors.WrapIf(err, "Error in registering repositories")
	}

	return nil
}
