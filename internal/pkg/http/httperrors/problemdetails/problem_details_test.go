package problemDetails

import (
	"net/http"
	"testing"

	customErrors "github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/http/httperrors/customerrors"

	"emperror.dev/errors"
	"github.com/stretchr/testify/assert"
)

func Test_Domain_Err(t *testing.T) {
	domainErr := NewDomainProblemDetail(http.StatusBadRequest, "Order with id '1' already completed", "stack")

	assert.Equal(t, "Order with id '1' already completed", domainErr.GetDetail())
	assert.Equal(t, "Domain Model Error", domainErr.GetTitle())
	assert.Equal(t, "stack", domainErr.GetStackTrace())
	assert.Equal(t, "https://httpstatuses.io/400", domainErr.GetType())
	assert.Equal(t, 400, domainErr.GetStatus())
}

func Test_Application_Err(t *testing.T) {
	applicationErr := NewApplicationProblemDetail(http.StatusBadRequest, "application_exceptions error", "stack")

	assert.Equal(t, "application_exceptions error", applicationErr.GetDetail())
	assert.Equal(t, "Application Service Error", applicationErr.GetTitle())
	assert.Equal(t, "stack", applicationErr.GetStackTrace())
	assert.Equal(t, "https://httpstatuses.io/400", applicationErr.GetType())
	assert.Equal(t, 400, applicationErr.GetStatus())
}

func Test_BadRequest_Err(t *testing.T) {
	badRequestError := NewBadRequestProblemDetail("bad-request error", "stack")

	assert.Equal(t, "bad-request error", badRequestError.GetDetail())
	assert.Equal(t, "Bad Request", badRequestError.GetTitle())
	assert.Equal(t, "stack", badRequestError.GetStackTrace())
	assert.Equal(t, "https://httpstatuses.io/400", badRequestError.GetType())
	assert.Equal(t, 400, badRequestError.GetStatus())
}

func Test_Parse_Error(t *testing.T) {
	// Bad-Request ProblemDetail
	badRequestError := errors.WrapIf(customErrors.NewBadRequestError("bad-request error"), "bad request error")
	badRequestPrb := ParseError(badRequestError)
	assert.NotNil(t, badRequestPrb)
	assert.Equal(t, badRequestPrb.GetStatus(), 400)

	// NotFound ProblemDetail
	notFoundError := customErrors.NewNotFoundError("notfound error")
	notfoundPrb := ParseError(notFoundError)
	assert.NotNil(t, notFoundError)
	assert.Equal(t, notfoundPrb.GetStatus(), 404)
}

func TestMap(t *testing.T) {
	Map[customErrors.BadRequestError](func(err customErrors.BadRequestError) ProblemDetailErr {
		return NewBadRequestProblemDetail(err.Message(), err.Error())
	})
	s := ResolveProblemDetail(customErrors.NewBadRequestError(""))
	_, ok := s.(ProblemDetailErr)
	assert.True(t, ok)
}
