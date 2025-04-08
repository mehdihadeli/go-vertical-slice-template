package integration

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/models"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	config2 "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/echoweb/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"
	defaultLogger "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger/defaultlogger"
	gormtestcontainer "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/test/containers/testcontainer/gorm"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/cockroachdb/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type IntegrationTestSharedFixture struct {
	EchoHttpOptions   *config2.EchoHttpOptions
	Log               logger.Logger
	ProductRepository contracts.ProductRepository
	GormDB            *gorm.DB
	BaseAddress       string
	Items             []*models.Product
	suite.Suite
}

func NewIntegrationTestSharedFixture(
	t *testing.T,
) *IntegrationTestSharedFixture {
	lifetimeCtx := context.Background()

	gormOptions, err := gormtestcontainer.GormContainerOptionsDecorator(
		t,
		lifetimeCtx,
		defaultLogger.GetLogger(),
	)
	if err != nil {
		t.Fatal(err)
	}

	// we can override test configuration for gorm with env with higher priority for test
	os.Setenv("GORM_OPTIONS_HOST", gormOptions.Host)
	os.Setenv("GORM_OPTIONS_PORT", strconv.Itoa(gormOptions.Port))
	os.Setenv("GORM_OPTIONS_USER", gormOptions.User)
	os.Setenv("GORM_OPTIONS_PASSWORD", gormOptions.Password)
	os.Setenv("GORM_OPTIONS_DATABASE_NAME", gormOptions.DBName)

	// this fix root working directory problem in our test environment inner our fixture
	environemnt.FixProjectRootWorkingDirectoryPath()

	application := app.NewTestApp().
		RunTest(t)

	integrationFixture := &IntegrationTestSharedFixture{
		EchoHttpOptions:   application.EchoOptions,
		Log:               application.Logger,
		ProductRepository: application.ProductRepository,
		GormDB:            application.GormDB,
		BaseAddress:       application.EchoOptions.BasePathAddress(),
	}

	return integrationFixture
}

func (i *IntegrationTestSharedFixture) SetupSuite() {
}

func (i *IntegrationTestSharedFixture) SetupTest() {
	i.Log.Info("SetupTest started")

	// migration will do in app configuration
	// seed data for our tests - app seed doesn't run in test environment
	res, err := seedDataManually(i.GormDB)
	if err != nil {
		i.Log.Error(errors.Wrap(err, "error in seeding data in postgres"))
	}

	i.Items = res
}

func (i *IntegrationTestSharedFixture) TearDownTest() {
	i.Log.Info("TearDownTest started")

	if err := i.cleanupPostgresData(); err != nil {
		i.Log.Error(errors.Wrap(err, "error in cleanup postgres data"))
	}
}

func (i *IntegrationTestSharedFixture) cleanupPostgresData() error {
	tables := []string{"products"}
	// Iterate over the tables and delete all records
	for _, table := range tables {
		err := i.GormDB.Exec("DELETE FROM " + table).Error

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
		return nil, errors.Wrap(err, "error in seed database")
	}

	return products, nil
}
