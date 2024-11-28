package controllers

import (
	"os"
	"time"

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

	accessSecret := os.Getenv("JWT_ACCESS_SECRET")
	accessExpiry := os.Getenv("JWT_ACCESS_EXPIRY")
	expiryDuration, err := time.ParseDuration(accessExpiry)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "invalid access token expiry duration", err)
	}

	accessToken, err := utils.CreateToken(user.ID.Hex(), accessSecret, expiryDuration)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "failed to create access token", err)
	}

	refreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	refreshExpiry := os.Getenv("JWT_REFRESH_EXPIRY")
	refreshExpiryDuration, err := time.ParseDuration(refreshExpiry)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "invalid refresh token expiry duration", err)
	}

	refreshToken, err := utils.CreateToken(user.ID.Hex(), refreshSecret, refreshExpiryDuration)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "failed to create refresh token", err)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(refreshExpiryDuration),
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	data := fiber.Map{
		"user":        user,
		"accessToken": accessToken,
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

	accessSecret := os.Getenv("JWT_ACCESS_SECRET")
	accessExpiry := os.Getenv("JWT_ACCESS_EXPIRY")
	expiryDuration, err := time.ParseDuration(accessExpiry)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "invalid access token expiry duration", err)
	}

	accessToken, err := utils.CreateToken(user.ID.Hex(), accessSecret, expiryDuration)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "failed to create access token", err)
	}

	refreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	refreshExpiry := os.Getenv("JWT_REFRESH_EXPIRY")
	refreshExpiryDuration, err := time.ParseDuration(refreshExpiry)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "invalid refresh token expiry duration", err)
	}

	refreshToken, err := utils.CreateToken(user.ID.Hex(), refreshSecret, refreshExpiryDuration)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "failed to create refresh token", err)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(refreshExpiryDuration),
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	data := fiber.Map{
		"user":        user,
		"accessToken": accessToken,
	}

	return utils.SendResponse(c, fiber.StatusOK, "login successful", data)
}
