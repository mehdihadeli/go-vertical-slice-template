//go:build integration
// +build integration

package creatingproduct

import (
	"context"
	"testing"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/commands"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creatingproduct/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/test/testfixtures/integration"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/stretchr/testify/suite"
)

type createProductIntegrationTests struct {
	*integration.IntegrationTestSharedFixture
}

func TestCreateProductIntegration(t *testing.T) {
	suite.Run(
		t,
		&createProductIntegrationTests{
			IntegrationTestSharedFixture: integration.NewIntegrationTestSharedFixture(t),
		},
	)
}

func (c *createProductIntegrationTests) Test_Should_Create_New_Product_To_DB() {
	ctx := context.Background()

	command := commands.NewCreateProductCommand(
		gofakeit.Name(),
		gofakeit.AdjectiveDescriptive(),
		gofakeit.Price(150, 6000),
	)

	result, err := mediatr.Send[*commands.CreateProductCommand, *dtos.CreateProductCommandResponse](
		ctx,
		command,
	)
	c.Require().NoError(err)

	c.Assert().NotNil(result)
	c.Assert().Equal(command.ProductID, result.ProductID)

	createdProduct, err := c.ProductRepository.GetProductById(
		ctx,
		result.ProductID,
	)
	c.Require().NoError(err)
	c.Assert().NotNil(createdProduct)
}

func (c *createProductIntegrationTests) Test_Should_Return_Error_For_Duplicate_Record() {
	ctx := context.Background()

	id := c.Items[0].ProductID

	command := &commands.CreateProductCommand{
		Name:        gofakeit.Name(),
		Description: gofakeit.AdjectiveDescriptive(),
		Price:       gofakeit.Price(150, 6000),
		ProductID:   id,
	}

	result, err := mediatr.Send[*commands.CreateProductCommand, *dtos.CreateProductCommandResponse](
		ctx,
		command,
	)
	c.Assert().Error(err)
	c.Assert().ErrorContains(err, "duplicate key value violates unique constraint")
	c.Assert().Nil(result)
}
