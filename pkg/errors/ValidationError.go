package errors

import (
	"khotba-online/pkg/validation"
	"net/http"
)

type ValidationError struct {
	Fields     []validation.ValidationError `json:"fields"`
	Message    string                        `json:"message"`
	StatusCode int                           `json:"status_code"`
}

func NewValidationError(fields []validation.ValidationError) *ValidationError {
	return &ValidationError{
		Fields:     fields,
		Message:    "Bad Request",
		StatusCode: http.StatusBadRequest,
	}
}

func (v ValidationError) Error() string {
	return v.Message
}
