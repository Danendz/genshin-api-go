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
	APIResponse struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Success bool        `json:"success"`
	}

	XValidator struct{
		validator *validator.Validate
	}

	ValidationError struct {
		Field string `json:"field"`
		Error string `json:"error"`
	}
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func NewApiResponse(message string, data interface{}, success bool) *APIResponse {
	return &APIResponse{
		Message: message,
		Data:    data,
		Success: success,
	}
}

func NewRestrictedApiRouter(router fiber.Router) fiber.Router {
	admin_username := os.Getenv("ADMIN_USERNAME")
	admin_password := os.Getenv("ADMIN_PASSWORD")

	restricted := router.Group("/", basicauth.New(
		basicauth.Config{
			Users: map[string]string{
				admin_username: admin_password,
			},
		},
	))

	return restricted
}

func (x XValidator) Validate(data interface{}) []*ValidationError{
	errors := []*ValidationError{}

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
