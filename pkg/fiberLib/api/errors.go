package api

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	error
	Code() uint
}

type Errors []Error

func (e Errors) MarshalJSON() ([]byte, error) {
	errors := make([]map[string]interface{}, 0, len(e))

	for _, err := range e {
		errors = append(errors, map[string]interface{}{
			"code":    err.Code(),
			"message": err.Error(),
		})
	}

	if len(errors) == 1 {
		return json.Marshal(errors[0])
	}

	return json.Marshal(errors)
}

type ParameterMissingError struct {
	ParameterName string
}

func (e ParameterMissingError) Error() string {
	return fmt.Sprintf("Mandatory parameter %q missing", e.ParameterName)
}

func (e ParameterMissingError) Code() uint {
	return ParameterMissingErrorCode
}

type MaxValueLimitError struct {
	Field string
	Limit string
}

func (e MaxValueLimitError) Error() string {
	return fmt.Sprintf("Max allowed values for parameter %q is %s", e.Field, e.Limit)
}

func (e MaxValueLimitError) Code() uint {
	return MaxValueLimitErrorCode
}

type InvalidBodyError struct{}

func (InvalidBodyError) Error() string {
	return "Invalid request body"
}

func (InvalidBodyError) Code() uint {
	return InvalidBodyErrorCode
}

type InvalidParameterTypeError struct {
	Parameter    string
	Type         string
	RequiredType string
}

func (i InvalidParameterTypeError) Error() string {
	return fmt.Sprintf(
		"Field %q value must be of type %q, %q given", i.Parameter, i.RequiredType, i.Type,
	)
}

func (i InvalidParameterTypeError) Code() uint {
	return InvalidParameterTypeErrorCode
}

type RedirectURINotAllowedError struct{}

func (RedirectURINotAllowedError) Error() string {
	return "Redirect URI is not allowed"
}

func (RedirectURINotAllowedError) Code() uint {
	return RedirectURINotAllowedErrorCode
}

type ActivitiesNotFoundError struct{}

type InvalidRequestPayloadError struct{}

func (InvalidRequestPayloadError) Error() string {
	return "Invalid request payload"
}

func (InvalidRequestPayloadError) Code() uint {
	return InvalidRequestErrorCode
}
