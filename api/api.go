package api

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/basicauth"
)


type APIResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Success bool `json:"success"`
}

func NewApiResponse(message string, data interface{}, success bool) *APIResponse {
	return &APIResponse{
		Message: message,
		Data: data,
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