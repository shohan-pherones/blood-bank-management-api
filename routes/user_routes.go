package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shohan-pherones/blood-bank-management.git/constants"
	"github.com/shohan-pherones/blood-bank-management.git/controllers"
	"github.com/shohan-pherones/blood-bank-management.git/middleware"
)

func RegisterUserRoutes(api fiber.Router) {
	user := api.Group("/users")

	user.Get("/", middleware.AuthMiddleware([]string{constants.RoleAdmin}), controllers.GetUsers)
	user.Get("/:id", middleware.AuthMiddleware([]string{constants.RoleUser, constants.RoleAdmin}), controllers.GetUser)
	user.Post("/register", controllers.RegisterUser)
	user.Post("/login", controllers.LoginUser)
}
