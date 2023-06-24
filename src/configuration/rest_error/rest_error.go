package rest_error

import (
	"net/http"
)

type RestError struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestError) Error() string {
	return r.Message
}

func NewRestError(message, error string, code int, causes []Causes) *RestError {
	return &RestError{
		Message: message,
		Err:     error,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "BAD_REQUEST",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestError {
	return &RestError{
		Message: message,
		Err:     "BAD_REQUEST_VALIDATION",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "FORBIDDEN",
		Code:    http.StatusForbidden,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "NOT_FOUND",
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "INTERNAL_SERVER_ERROR",
		Code:    http.StatusInternalServerError,
	}
}
