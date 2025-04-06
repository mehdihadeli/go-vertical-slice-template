package getting_product_by_id

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/mehdihadeli/go-mediatr"
	dtos2 "github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/queries"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/mappings"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/models"
	"github.com/mehdihadeli/go-vertical-slice-template/test/testfixtures/unittest"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type getProductByIdHandlerUnitTests struct {
	*unittest.UnitTestSharedFixture
	handler mediatr.RequestHandler[*queries.GetProductByIdQuery, *dtos2.GetProductByIdQueryResponse]
}

func TestCreateProductHandlerUnit(t *testing.T) {
	suite.Run(t, &getProductByIdHandlerUnitTests{
		UnitTestSharedFixture: unittest.NewUnitTestSharedFixture(t),
	},
	)
}

func (c *getProductByIdHandlerUnitTests) SetupTest() {
	// call base SetupTest hook before running child hook
	c.UnitTestSharedFixture.SetupTest()
	c.handler = queries.NewGetProductByIdHandler(c.ProductRepository)
}

func (c *getProductByIdHandlerUnitTests) TearDownTest() {
	// call base TearDownTest hook before running child hook
	c.UnitTestSharedFixture.TearDownTest()
}

func (g *getProductByIdHandlerUnitTests) Test_Handle_Should_Return_Product_When_Exists() {
	// Arrange
	productID := uuid.NewV4()
	expectedProduct := &models.Product{
		ProductID:   productID,
		Name:        gofakeit.ProductName(),
		Description: gofakeit.Sentence(10),
		Price:       gofakeit.Price(10, 1000),
		CreatedAt:   gofakeit.DateRange(time.Now().AddDate(0, -1, 0), time.Now()), // Random date within last year
		UpdatedAt:   gofakeit.DateRange(time.Now().AddDate(0, -1, 0), time.Now()), // Random date within last year
	}

	g.ProductRepository.On("GetProductById", g.Ctx, productID).
		Return(expectedProduct, nil).
		Once()

	query := &queries.GetProductByIdQuery{ProductID: productID}

	// Act
	response, err := g.handler.Handle(g.Ctx, query)

	// Assert
	g.Require().NoError(err)
	g.Require().NotNil(response)
	g.Require().NotNil(response.Product)
	g.Equal(mappings.MapProductToProductDto(expectedProduct), response.Product)
	g.ProductRepository.AssertCalled(g.T(), "GetProductById", g.Ctx, productID)
	g.ProductRepository.AssertNumberOfCalls(g.T(), "GetProductById", 1)
}

func (g *getProductByIdHandlerUnitTests) Test_Handle_Should_Return_Error_When_Product_Not_Found() {
	// Arrange
	productID := uuid.NewV4()
	expectedError := errors.New("product not found")

	g.ProductRepository.On("GetProductById", g.Ctx, productID).
		Return(nil, expectedError).
		Once()

	query := &queries.GetProductByIdQuery{ProductID: productID}

	// Act
	response, err := g.handler.Handle(g.Ctx, query)

	// Assert
	g.Require().Error(err)
	g.Require().Nil(response)
	g.Contains(err.Error(), fmt.Sprintf("product with id %s not found", productID))
	g.ProductRepository.AssertCalled(g.T(), "GetProductById", g.Ctx, productID)
	g.ProductRepository.AssertNumberOfCalls(g.T(), "GetProductById", 1)
}

func (g *getProductByIdHandlerUnitTests) Test_Handle_Should_Return_Error_When_Query_Is_Nil() {
	// Act
	response, err := g.handler.Handle(g.Ctx, nil)

	// Assert
	g.Require().Error(err)
	g.Require().Nil(response)
	g.Equal("query cannot be nil", err.Error())
}

func (g *getProductByIdHandlerUnitTests) Test_Handle_Should_Return_Error_When_ProductID_Is_Zero() {
	// Arrange
	query := queries.NewGetProductByIdQuery(uuid.Nil)

	// Act
	response, err := g.handler.Handle(g.Ctx, query)

	// Assert
	g.Require().Error(err)
	g.Require().Nil(response)
	g.Contains(err.Error(), "product id cannot be empty")
}

func (g *getProductByIdHandlerUnitTests) Test_Handle_Should_Return_Error_When_Repository_Returns_Nil_Product_Without_Error() {
	// Arrange
	productID := uuid.NewV4()

	g.ProductRepository.On("GetProductById", g.Ctx, productID).
		Return(nil, errors.New("product not found")).
		Once()

	query := queries.NewGetProductByIdQuery(productID)

	// Act
	response, err := g.handler.Handle(g.Ctx, query)

	// Assert
	g.Require().NotNil(err)
	g.Require().Nil(response)
	g.Contains(err.Error(), fmt.Sprintf("product with id %s not found", productID))
}
