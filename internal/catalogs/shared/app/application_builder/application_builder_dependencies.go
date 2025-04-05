package applicationbuilder

import (
	"github.com/mehdihadeli/go-vertical-slice-template/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/echoweb"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"
	"gorm.io/gorm"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts/params"
	endpoints2 "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creating_product/endpoints"
	endpoints3 "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/getting_product_by_id/endpoints"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/repository"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/database"
)

func (b *ApplicationBuilder) AddInfrastructure() {
	err := config.AddAppConfig(b.Container)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = database.AddGorm(b.Container)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = echoweb.AddEcho(b.Container)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddRoutes() {
	err := b.Container.Provide(func(productRouteParams *params.ProductRouteParams) ([]contracts.Endpoint, error) {
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
	err := b.Container.Provide(func(g *gorm.DB, l logger.Logger) (contracts.ProductRepository, error) {
		return repository.NewInMemoryProductRepository(g, l), nil
	})

	if err != nil {
		b.Logger.Fatal(err)
	}
}
