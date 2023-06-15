package applicationbuilder

import (
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/go-vertical-slice-template/internal/products/api"
	"github.com/go-vertical-slice-template/internal/products/repository"
	"github.com/go-vertical-slice-template/internal/shared/database"
)

func (b *ApplicationBuilder) AddEcho() {
	echoDep := di.Def{
		Name:  "echo",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return echo.New(), nil
		},
	}

	err := b.Services.Add(echoDep)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddGorm() {
	gormDep := di.Def{
		Name:  "gorm",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return database.NewGormDB()
		},
	}

	err := b.Services.Add(gormDep)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddLogger() {
	logDep := di.Def{
		Name:  "zap",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return b.Logger, nil
		},
	}

	err := b.Services.Add(logDep)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddControllers() {
	productControllerDep := di.Def{
		Name:  "productController",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			e := ctn.Get("echo").(*echo.Echo)
			_ = ctn.Get("zap").(*zap.SugaredLogger)
			return api.NewProductsController(e), nil
		},
	}

	err := b.Services.Add(productControllerDep)
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
