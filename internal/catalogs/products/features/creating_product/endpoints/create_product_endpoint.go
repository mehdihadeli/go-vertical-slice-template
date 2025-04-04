package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts/params"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creating_product/commands"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/features/creating_product/dtos"
)

type createProductEndpoint struct {
	*params.ProductRouteParams
}

func NewCreteProductEndpoint(params *params.ProductRouteParams) contracts.Endpoint {
	return &createProductEndpoint{ProductRouteParams: params}
}

func (ep *createProductEndpoint) MapEndpoint() {
	ep.ProductsGroup.POST("", ep.handler())
}

// CreateProduct
// @Tags Products
// @Summary Create product
// @Description Create new product item
// @Accept json
// @Produce json
// @Param CreateProductRequestDto body creatingProductsDtos.CreateProductRequestDto true "Product data"
// @Success 201 {object} creatingProductsDtos.CreateProductResponseDto
// @Router /api/v1/products [post]
func (ep *createProductEndpoint) handler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &dtos.CreateProductRequestDto{}
		if err := ctx.Bind(request); err != nil {
			return err
		}

		if err := ep.Validator.StructCtx(ctx.Request().Context(), request); err != nil {
			return err
		}

		command := commands.NewCreateProductCommand(request.Name, request.Description, request.Price)
		result, err := mediatr.Send[*commands.CreateProductCommand, *dtos.CreateProductCommandResponse](ctx.Request().Context(), command)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, result)
	}
}
