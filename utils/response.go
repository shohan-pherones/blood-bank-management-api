package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SendResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	response := Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	return c.Status(status).JSON(response)
}

func SendError(c *fiber.Ctx, status int, message string, err error) error {
	response := Response{
		Status:  "error",
		Message: message,
		Error:   err.Error(),
	}

	return c.Status(status).JSON(response)
}
