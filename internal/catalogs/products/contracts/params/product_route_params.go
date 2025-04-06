package params

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ProductRouteParams struct {
	Logger        logger.Logger
	ProductsGroup *echo.Group
	Validator     *validator.Validate
}
