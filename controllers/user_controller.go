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

	data := fiber.Map{
		"user": user,
	}

	return utils.SendResponse(c, fiber.StatusCreated, "user registered successfully", data)
}

func LoginUser(c *fiber.Ctx) error {
	type LoginPayload struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	var payload LoginPayload
	if err := c.BodyParser(&payload); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "failed to parse request body", err)
	}

	userService := services.UserService{}
	user, err := userService.LoginUser(payload.Email, payload.Password)
	if err != nil {
		return utils.SendError(c, fiber.StatusUnauthorized, "login failed", err)
	}

	data := fiber.Map{
		"user": user,
	}

	return utils.SendResponse(c, fiber.StatusOK, "login successful", data)
}
