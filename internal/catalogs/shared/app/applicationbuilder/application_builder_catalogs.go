package applicationbuilder

import (
	"net/http"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts/params"
	creatingproductendpoint "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/endpoints"
	gettingproductbyidendpoint "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/endpoints"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/repository"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
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
	b.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Catalogs Api!")
	})

	// https://echo.labstack.com/docs/routing
	v1 := b.echo.Group("/api/v1")
	productsEchoGroup := v1.Group("/products")
	productsRouteParams := &params.ProductRouteParams{
		Logger:        b.Logger,
		Validator:     validator.New(),
		ProductsGroup: productsEchoGroup,
	}

	createProductEndpoint := creatingproductendpoint.NewCreteProductEndpoint(productsRouteParams)
	getProductById := gettingproductbyidendpoint.NewGetProductByIdEndpoint(productsRouteParams)
	endpoints := []contracts.Endpoint{createProductEndpoint, getProductById}

	b.endpoints = endpoints

	return nil
}

func (b *ApplicationBuilder) addRepositories() error {
	b.productRepository = repository.NewProductRepositoryGorm(b.gormDB, b.Logger)

	return nil
}
