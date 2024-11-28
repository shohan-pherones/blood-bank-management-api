package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shohan-pherones/blood-bank-management.git/controllers"
)

func RegisterUserRoutes(api fiber.Router) {
	user := api.Group("/users")

	user.Post("/register", controllers.RegisterUser)
	user.Post("/login", controllers.LoginUser)
}
