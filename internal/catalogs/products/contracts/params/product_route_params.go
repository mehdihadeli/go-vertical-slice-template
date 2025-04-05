package params

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"
)

type ProductRouteParams struct {
	Logger        logger.Logger
	ProductsGroup *echo.Group
	Validator     *validator.Validate
}
