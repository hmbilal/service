package api

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"net/http"
	"strings"

	"github.com/andreiavrammsd/validator"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	//if ctx.Response() {
	//	return
	//}

	if handleErr := doHandling(err, ctx); handleErr != nil {
		return handleErr
	}

	return nil
}

func doHandling(err error, ctx *fiber.Ctx) error {
	if errs := formatValidationErrors(err); errs != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": errs,
		})
	}

	var httpErr *fiber.Error

	if ok := errors.As(err, &httpErr); ok {
		log.Warn(httpErr.Message)

		return ctx.SendStatus(httpErr.Code)
	}

	log.Error(err)

	return ctx.SendStatus(http.StatusInternalServerError)
}

func formatValidationErrors(err error) Errors {
	var (
		apiErr                 Error
		validationErrs         validator.ValidationErrors
		jsonSyntaxError        *json.SyntaxError
		jsonUnmarshalTypeError *json.UnmarshalTypeError
	)

	if errors.As(err, &apiErr) {
		return Errors{apiErr}
	}

	if errors.As(err, &validationErrs) {
		formattedErrs := make(Errors, 0)

		for _, validationErr := range validationErrs {
			if formattedErr := resolveValidationError(validationErr); formattedErr != nil {
				formattedErrs = append(formattedErrs, formattedErr)
			}
		}

		return formattedErrs
	}

	if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
		return Errors{InvalidBodyError{}}
	}

	if errors.As(err, &jsonSyntaxError) {
		return Errors{InvalidBodyError{}}
	} else if errors.As(err, &jsonUnmarshalTypeError) {
		return Errors{InvalidParameterTypeError{
			Parameter:    jsonUnmarshalTypeError.Field,
			Type:         jsonUnmarshalTypeError.Type.String(),
			RequiredType: jsonUnmarshalTypeError.Value,
		}}
	}

	return nil
}

func resolveValidationError(err validator.FieldError) Error {
	field := removeNamespaceRoot(err.Namespace())

	switch err.Tag() {
	case "required", "notblank":
		return ParameterMissingError{ParameterName: field}
	case "max":
		return MaxValueLimitError{Field: field, Limit: err.Param()}
	}

	return nil
}

func removeNamespaceRoot(namespace string) string {
	namespaceAsBytes := []byte(namespace)
	separatorIndex := strings.Index(namespace, ".")

	return strings.ToLower(string(namespaceAsBytes[separatorIndex+1:]))
}
