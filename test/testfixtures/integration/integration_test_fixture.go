package integration

import (
	"context"
	"testing"
	"time"

	"github.com/mehdihadeli/go-vertical-slice-template/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/models"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"
	gotmtestcontainer "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/test/containers/testcontainer/gorm"

	"emperror.dev/errors"
	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/dig"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type IntegrationTestSharedFixture struct {
	Cfg               *config.Config
	Log               logger.Logger
	Container         *dig.Container
	ProductRepository contracts.ProductRepository
	Gorm              *gorm.DB
	BaseAddress       string
	Items             []*models.Product
	suite.Suite
}

func NewIntegrationTestSharedFixture(
	t *testing.T,
) *IntegrationTestSharedFixture {
	// this fix root working directory problem in our test environment inner our fixture
	environemnt.FixProjectRootWorkingDirectoryPath()

	lifetimeCtx := context.Background()
	container := app.NewTestApp().
		WithOverrideBuilder(gotmtestcontainer.GormContainerOptionsDecorator(t, lifetimeCtx)).
		RunTest(t)

	integrationFixture := &IntegrationTestSharedFixture{}

	err := container.Invoke(
		func(l logger.Logger, db *gorm.DB, cfg *config.Config, productRepository contracts.ProductRepository) {
			integrationFixture.Log = l
			integrationFixture.Container = container
			integrationFixture.Gorm = db
			integrationFixture.Cfg = cfg
			integrationFixture.BaseAddress = cfg.EchoHttpOptions.BasePathAddress()
			integrationFixture.ProductRepository = productRepository
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	return integrationFixture
}

func (i *IntegrationTestSharedFixture) SetupSuite() {
}

func (i *IntegrationTestSharedFixture) SetupTest() {
	i.Log.Info("SetupTest started")

	// migration will do in app configuration
	// seed data for our tests - app seed doesn't run in test environment
	res, err := seedDataManually(i.Gorm)
	if err != nil {
		i.Log.Error(errors.WrapIf(err, "error in seeding data in postgres"))
	}

	i.Items = res
}

func (i *IntegrationTestSharedFixture) TearDownTest() {
	i.Log.Info("TearDownTest started")

	if err := i.cleanupPostgresData(); err != nil {
		i.Log.Error(errors.WrapIf(err, "error in cleanup postgres data"))
	}
}

func (i *IntegrationTestSharedFixture) cleanupPostgresData() error {
	tables := []string{"products"}
	// Iterate over the tables and delete all records
	for _, table := range tables {
		err := i.Gorm.Exec("DELETE FROM " + table).Error

		return err
	}

	return nil
}

func seedDataManually(gormDB *gorm.DB) ([]*models.Product, error) {
	products := []*models.Product{
		{
			ProductID:   uuid.NewV4(),
			Name:        gofakeit.Name(),
			CreatedAt:   time.Now(),
			Description: gofakeit.AdjectiveDescriptive(),
			Price:       gofakeit.Price(100, 1000),
		},
		{
			ProductID:   uuid.NewV4(),
			Name:        gofakeit.Name(),
			CreatedAt:   time.Now(),
			Description: gofakeit.AdjectiveDescriptive(),
			Price:       gofakeit.Price(100, 1000),
		},
	}

	err := gormDB.CreateInBatches(products, len(products)).Error
	if err != nil {
		return nil, errors.WrapIf(err, "error in seed database")
	}

	return products, nil
}
