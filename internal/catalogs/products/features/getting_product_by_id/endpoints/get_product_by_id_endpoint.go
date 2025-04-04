package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts/params"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/getting_product_by_id/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/getting_product_by_id/queries"
)

type getProductByIdEndpoint struct {
	*params.ProductRouteParams
}

func NewGetProductByIdEndpoint(params *params.ProductRouteParams) contracts.Endpoint {
	return &getProductByIdEndpoint{ProductRouteParams: params}
}

func (ep *getProductByIdEndpoint) MapEndpoint() {
	ep.ProductsGroup.GET("/:id", ep.handler())
}

// GetProductByID
// @Tags Products
// @Summary Get product
// @Description Get product by id
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} gettingProductByIdDtos.GetProductByIdResponseDto
// @Router /api/v1/products/{id} [get]
func (ep *getProductByIdEndpoint) handler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &dtos.GetProductByIdRequestDto{}
		if err := ctx.Bind(request); err != nil {
			return err
		}

		query := queries.NewGetProductByIdQuery(request.ProductId)

		if err := ep.Validator.StructCtx(ctx.Request().Context(), query); err != nil {
			return err
		}

		queryResult, err := mediatr.Send[*queries.GetProductByIdQuery, *dtos.GetProductByIdQueryResponse](ctx.Request().Context(), query)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, queryResult)
	}
}
