//go:build e2e
// +build e2e

package creatingproduct

import (
	"context"
	"net/http"
	"testing"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/test/testfixtures/integration"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/suite"
)

type createProductE2ETest struct {
	*integration.IntegrationTestSharedFixture
}

func TestCreateProductE2E(t *testing.T) {
	suite.Run(
		t,
		&createProductE2ETest{
			IntegrationTestSharedFixture: integration.NewIntegrationTestSharedFixture(t),
		},
	)
}

func (c *createProductE2ETest) Test_Should_Return_Created_Status_With_Valid_Input() {
	request := dtos.CreateProductRequestDto{
		Description: gofakeit.AdjectiveDescriptive(),
		Price:       gofakeit.Price(100, 1000),
		Name:        gofakeit.Name(),
	}

	// create httpexpect instance
	expect := httpexpect.Default(c.T(), c.BaseAddress)

	expect.POST("/api/v1/products").
		WithContext(context.Background()).
		WithJSON(request).
		Expect().
		Status(http.StatusCreated)
}

// Input validations
func (c *createProductE2ETest) Test_Should_Return_Bad_Request_Status_With_Invalid_Price_Input() {
	request := dtos.CreateProductRequestDto{
		Description: gofakeit.AdjectiveDescriptive(),
		Price:       0,
		Name:        gofakeit.Name(),
	}

	// create httpexpect instance
	expect := httpexpect.Default(c.T(), c.BaseAddress)

	expect.POST("/api/v1/products").
		WithContext(context.Background()).
		WithJSON(request).
		Expect().
		Status(http.StatusBadRequest)
}
