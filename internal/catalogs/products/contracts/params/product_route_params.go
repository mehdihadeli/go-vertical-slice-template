package params

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ProductRouteParams struct {
	Logger        *zap.SugaredLogger
	ProductsGroup *echo.Group
	Validator     *validator.Validate
}
