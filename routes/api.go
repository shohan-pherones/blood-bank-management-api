package routes

import (
	"github.com/gofiber/fiber/v2"
)

func MainAPIEndpoint(app *fiber.App) fiber.Router {
	return app.Group("/api/v1")
}
