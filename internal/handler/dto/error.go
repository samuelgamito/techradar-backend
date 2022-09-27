package dto

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const (
	defaultErrorMessage           = "Not able to process the request, contact the admin"
	maxErrorMessage               = "%s max values is %s, actual value is %d."
	minErrorMessage               = "%s min values is %s, actual value is %d."
	missingRequiredFieldMessage   = "Field %s is required."
	defaultValidationErrorMessage = "Validation fails on %s"
)

type (
	ErrorResponse struct {
		StatusCode int
		Body       ErrorBodyDTO
	}

	ErrorBodyDTO struct {
		Messages []string `json:"messages"`
	}
)

func (m *ErrorResponse) Error() string {
	return "has validation error"
}

func DefaultError() *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 500,
		Body: ErrorBodyDTO{
			Messages: []string{defaultErrorMessage},
		},
	}
}

func buildErrorMessage(e validator.FieldError) string {
	switch errType := e.Tag(); errType {
	case "min":
		return fmt.Sprintf(minErrorMessage, e.Field(), e.Param(), e.Value())
	case "max":
		return fmt.Sprintf(maxErrorMessage, e.Field(), e.Param(), e.Value())
	case "required":
		return fmt.Sprintf(missingRequiredFieldMessage, e.Field())
	default:
		return fmt.Sprintf(defaultValidationErrorMessage, e.Field())
	}
}
