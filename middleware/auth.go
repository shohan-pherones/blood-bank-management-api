package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shohan-pherones/blood-bank-management.git/utils"
)

func AuthMiddleware(expectedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.SendError(c, fiber.StatusUnauthorized, "authorization header is missing", nil)
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return utils.SendError(c, fiber.StatusUnauthorized, "invalid token format", nil)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		accessSecret := os.Getenv("JWT_ACCESS_SECRET")
		claims, err := utils.VerifyToken(tokenString, accessSecret)
		if err != nil {
			return utils.SendError(c, fiber.StatusUnauthorized, "invalid or expired access token", err)
		}

		log.Print(claims)

		userRole, ok := claims["role"]
		if !ok {
			return utils.SendError(c, fiber.StatusForbidden, "invalid token: role missing", nil)
		}

		roleStr, ok := userRole.(string)
		if !ok {
			return utils.SendError(c, fiber.StatusForbidden, "invalid token: role is invalid", nil)
		}

		roleMatch := false
		for _, role := range expectedRoles {
			if roleStr == role {
				roleMatch = true
				break
			}
		}

		if !roleMatch {
			return utils.SendError(c, fiber.StatusForbidden, "you do not have the required permissions", nil)
		}

		c.Locals("userID", claims["user_id"])
		c.Locals("role", roleStr)

		return c.Next()
	}
}
