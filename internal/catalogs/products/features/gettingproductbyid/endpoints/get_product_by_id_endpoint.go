package endpoints

import (
	"net/http"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts/params"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/dtos"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/gettingproductbyid/queries"
	customErrors "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/httperrors/customerrors"

	"emperror.dev/errors"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
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
			return customErrors.NewBadRequestErrorWrap(
				err,
				"error in the binding request",
			)
		}

		query := queries.NewGetProductByIdQuery(request.ProductId)

		if err := ep.Validator.StructCtx(ctx.Request().Context(), query); err != nil {
			return customErrors.NewValidationErrorWrap(err, "validation error")
		}

		queryResult, err := mediatr.Send[*queries.GetProductByIdQuery, *dtos.GetProductByIdQueryResponse](
			ctx.Request().Context(),
			query,
		)
		if err != nil {
			return errors.WithMessage(
				err,
				"error in sending GetProductByIdQuery",
			)
		}

		return ctx.JSON(http.StatusOK, queryResult)
	}
}
