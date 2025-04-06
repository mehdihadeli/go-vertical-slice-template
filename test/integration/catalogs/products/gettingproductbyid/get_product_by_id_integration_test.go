//go:build integration
// +build integration

package gettingproductbyid

import (
	"context"
	"fmt"
	"testing"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/queries"
	"github.com/mehdihadeli/go-vertical-slice-template/test/testfixtures/integration"

	"github.com/mehdihadeli/go-mediatr"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
)

type getProductByIdIntegrationTests struct {
	*integration.IntegrationTestSharedFixture
}

func TestGetProductByIdIntegration(t *testing.T) {
	suite.Run(
		t,
		&getProductByIdIntegrationTests{
			IntegrationTestSharedFixture: integration.NewIntegrationTestSharedFixture(t),
		},
	)
}

func (c *getProductByIdIntegrationTests) Test_Should_Returns_Existing_Product_From_DB_With_Correct_Properties() {
	ctx := context.Background()

	id := c.Items[0].ProductID
	query := queries.NewGetProductByIdQuery(id)
	result, err := mediatr.Send[*queries.GetProductByIdQuery, *dtos.GetProductByIdQueryResponse](
		ctx,
		query,
	)

	c.Require().NoError(err)
	c.NotNil(result)
	c.NotNil(result.Product)
	c.Equal(id, result.Product.ProductID)
}

func (c *getProductByIdIntegrationTests) Test_Should_Returns_NotFound_Error_When_Record_DoesNot_Exists() {
	ctx := context.Background()

	id := uuid.NewV4()
	query := queries.NewGetProductByIdQuery(id)

	result, err := mediatr.Send[*queries.GetProductByIdQuery, *dtos.GetProductByIdQueryResponse](
		ctx,
		query,
	)

	c.Require().Error(err)
	c.Require().ErrorContains(err, fmt.Sprintf("product with id %s not found", id))
	c.Nil(result)
}
