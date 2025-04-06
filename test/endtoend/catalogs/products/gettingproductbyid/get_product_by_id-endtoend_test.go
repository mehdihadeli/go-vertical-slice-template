package gettingproductbyid

import (
	"context"
	"net/http"
	"testing"

	"github.com/mehdihadeli/go-vertical-slice-template/test/testfixtures/integration"

	"github.com/gavv/httpexpect/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
)

type getProductByIdE2ETest struct {
	*integration.IntegrationTestSharedFixture
}

func TestGetProductByIdE2E(t *testing.T) {
	suite.Run(
		t,
		&getProductByIdE2ETest{
			IntegrationTestSharedFixture: integration.NewIntegrationTestSharedFixture(t),
		},
	)
}

func (c *getProductByIdE2ETest) Test_Should_Return_Ok_Status_With_Valid_Id() {
	ctx := context.Background()

	// create httpexpect instance
	expect := httpexpect.Default(c.T(), c.BaseAddress)

	id := c.Items[0].ProductID

	expect.GET("/api/v1/products/{id}").
		WithPath("id", id.String()).
		WithContext(ctx).
		Expect().
		Status(http.StatusOK)
}

// Input validations
func (c *getProductByIdE2ETest) Test_Should_Return_NotFound_Status_With_Invalid_Id() {
	ctx := context.Background()

	// create httpexpect instance
	expect := httpexpect.Default(c.T(), c.BaseAddress)

	id := uuid.NewV4()

	expect.GET("/api/v1/products/{id}").
		WithPath("id", id.String()).
		WithContext(ctx).
		Expect().
		Status(http.StatusNotFound)
}
