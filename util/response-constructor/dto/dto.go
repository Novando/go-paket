package dto

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type StandardResponse struct {
	Value   any    `json:"value"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

type ErrorValue struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type MetaNormal struct {
	TotalData int `json:"totalData"`
	TotalPage int `json:"totalPage"`
	Page      int `json:"page"`
	Size      int `json:"size"`
}

type MetaCursor struct {
	Size int    `json:"size"`
	Next string `json:"next"`
	Prev string `json:"prev"`
}

// ArrayResponse represents a standard array response structure
type ArrayResponse[T any] struct {
	Data []T `json:"data"`
}

// ArrayResponseWithMetaNormal represents an array response with MetaNormal metadata
type ArrayResponseWithMetaNormal[T any] struct {
	Data []T        `json:"data"`
	Meta MetaNormal `json:"meta"`
}

// ArrayResponseWithMetaCursor represents an array response with MetaCursor metadata
type ArrayResponseWithMetaCursor[T any] struct {
	Data []T        `json:"data"`
	Meta MetaCursor `json:"meta"`
}

func (sr StandardResponse) Error() string {
	return sr.Message
}

// NewErrorResponse creates a new StandardResponse with an error
func NewErrorResponse(err error) StandardResponse {
	var valErrs validator.ValidationErrors
	if errors.As(err, &valErrs) {
		var errVals []ErrorValue
		for _, e := range valErrs {
			errVals = append(errVals, ErrorValue{Field: e.Field(), Error: e.Error()})
		}
		return StandardResponse{Code: "VALIDATION_ERROR", Message: "Some of the data sent were invalid", Value: errVals}
	}

	var stdErr StandardResponse
	if errors.As(err, &stdErr) {
		return StandardResponse{Code: stdErr.Code, Message: stdErr.Message, Value: stdErr.Value}
	}

	return StandardResponse{Code: "UNKNOWN_ERROR", Message: "Something went wrong"}
}

func NewNotImplementedError() StandardResponse {
	return StandardResponse{Code: "NOT_IMPLEMENTED_YET", Message: "The function is still on progress", Value: nil}
}
