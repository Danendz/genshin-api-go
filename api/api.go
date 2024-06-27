package api

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/basicauth"
)

type (
	Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Success bool        `json:"success"`
	}

	XValidator struct {
		validator *validator.Validate
	}

	ValidationError struct {
		Field string `json:"field"`
		Error string `json:"error"`
	}
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func NewApiResponse(message string, data interface{}, success bool) *Response {
	return &Response{
		Message: message,
		Data:    data,
		Success: success,
	}
}

func NewRestrictedApiRouter(router fiber.Router) fiber.Router {
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	restricted := router.Group("/", basicauth.New(
		basicauth.Config{
			Users: map[string]string{
				adminUsername: adminPassword,
			},
		},
	))

	return restricted
}

func (x XValidator) Validate(data interface{}) []*ValidationError {
	var errors []*ValidationError

	if err := x.validator.Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ValidationError{
				Field: err.Field(),
				Error: validationMessage(err.Tag(), err.Param()),
			})
		}
	}

	return errors
}

func validationMessage(tag string, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "min":
		return fmt.Sprintf("This field must be greater than or equal to %v", param)
	case "max":
		return fmt.Sprintf("This field must be less than or equal to %v", param)
	default:
		return "This field is invalid"
	}
}

func NewValidator() *XValidator {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return &XValidator{
		validator: validate,
	}
}
