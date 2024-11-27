package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shohan-pherones/blood-bank-management.git/models"
	"github.com/shohan-pherones/blood-bank-management.git/services"
	"github.com/shohan-pherones/blood-bank-management.git/utils"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.UserModel
	if err := c.BodyParser(&user); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "failed to parse request body", err)
	}

	userService := services.UserService{}
	if err := userService.RegisterUser(&user); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "failed to register user", err)
	}

	return utils.SendResponse(c, fiber.StatusCreated, "user registered successfully", user)
}
