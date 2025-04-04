package applicationbuilder

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/mehdihadeli/go-vertical-slice-template/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts/params"
	endpoints2 "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creating_product/endpoints"
	endpoints3 "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/getting_product_by_id/endpoints"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/repository"
	config2 "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/database"
)

func (b *ApplicationBuilder) AddCore() {

	logDep := di.Def{
		Name:  "zap",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return b.Logger, nil
		},
	}

	configDep := di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			env := ctn.Get("env").(environemnt.Environment)
			return config.NewConfig(env)
		},
	}

	err := config2.AddEnv(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = b.Services.Add(logDep)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = b.Services.Add(configDep)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddInfrastructure() {
	err := database.AddGorm(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = addEcho(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddRoutes() {
	routesDep := di.Def{
		Name:  "routes",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			productRouteParams := ctn.Get("productRouteGroup").(*params.ProductRouteParams)

			createProductEndpoint := endpoints2.NewCreteProductEndpoint(productRouteParams)
			getProductById := endpoints3.NewGetProductByIdEndpoint(productRouteParams)
			endpoints := []contracts.Endpoint{createProductEndpoint, getProductById}

			return endpoints, nil
		},
	}

	err := b.Services.Add(routesDep)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddRepositories() {
	productRepositoryDep := di.Def{
		Name:  "productRepository",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			g := ctn.Get("gorm").(*gorm.DB)
			l := ctn.Get("zap").(*zap.SugaredLogger)
			return repository.NewInMemoryProductRepository(g, l), nil
		},
	}

	err := b.Services.Add(productRepositoryDep)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func addEcho(container *di.Builder) error {
	echoDep := di.Def{
		Name:  "echo",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return echo.New(), nil
		},
	}

	productGroupDep := di.Def{
		Name:  "productRouteGroup",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			echo := ctn.Get("echo").(*echo.Echo)
			logger := ctn.Get("zap").(*zap.SugaredLogger)

			v1 := echo.Group("/api/v1")
			products := v1.Group("/products")

			productsRouteParams := &params.ProductRouteParams{
				Logger:        logger,
				Validator:     validator.New(),
				ProductsGroup: products,
			}

			return productsRouteParams, nil
		},
	}

	err := container.Add(echoDep)
	if err != nil {
		return err
	}

	err = container.Add(productGroupDep)
	if err != nil {
		return err
	}

	return nil
}
