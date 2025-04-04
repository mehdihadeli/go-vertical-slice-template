package applicationbuilder

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/echoweb"
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
	err := b.Services.Provide(func() *zap.SugaredLogger {
		return b.Logger
	})
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = b.Services.Provide(func(environment environemnt.Environment) (*config.Config, error) {
		return config.NewConfig(environment)
	})
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = config2.AddEnv(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddInfrastructure() {
	err := database.AddGorm(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = echoweb.AddEcho(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddRoutes() {
	err := b.Services.Provide(func(productRouteParams *params.ProductRouteParams) ([]contracts.Endpoint, error) {
		createProductEndpoint := endpoints2.NewCreteProductEndpoint(productRouteParams)
		getProductById := endpoints3.NewGetProductByIdEndpoint(productRouteParams)
		endpoints := []contracts.Endpoint{createProductEndpoint, getProductById}

		return endpoints, nil
	})

	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddRepositories() {
	err := b.Services.Provide(func(g *gorm.DB, l *zap.SugaredLogger) (contracts.ProductRepository, error) {
		return repository.NewInMemoryProductRepository(g, l), nil
	})

	if err != nil {
		b.Logger.Fatal(err)
	}
}
