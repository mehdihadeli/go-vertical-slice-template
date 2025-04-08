package applicationbuilder

import (
	"net/http"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts/params"
	creatingproductendpoint "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/endpoints"
	gettingproductbyidendpoint "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/endpoints"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/repository"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/dependency"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
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
	dependency.Add[[]contracts.Endpoint](
		b.ServiceCollection,
		func(sp *dependency.ServiceProvider) ([]contracts.Endpoint, error) {
			e := dependency.GetGenericRequiredService[*echo.Echo](sp)

			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, "Catalogs Api!")
			})

			// https://echo.labstack.com/docs/routing
			v1 := e.Group("/api/v1")
			products := v1.Group("/products")
			productsRouteParams := &params.ProductRouteParams{
				Logger:        b.Logger,
				Validator:     validator.New(),
				ProductsGroup: products,
			}

			createProductEndpoint := creatingproductendpoint.NewCreteProductEndpoint(productsRouteParams)
			getProductById := gettingproductbyidendpoint.NewGetProductByIdEndpoint(productsRouteParams)
			endpoints := []contracts.Endpoint{createProductEndpoint, getProductById}

			return endpoints, nil
		},
	)

	return nil
}

func (b *ApplicationBuilder) addRepositories() error {
	dependency.Add[contracts.ProductRepository](
		b.ServiceCollection,
		func(sp *dependency.ServiceProvider) (contracts.ProductRepository, error) {
			gormDB := dependency.GetGenericRequiredService[*gorm.DB](sp)

			return repository.NewProductRepositoryGorm(gormDB, b.Logger), nil
		},
	)

	return nil
}
