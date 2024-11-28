package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shohan-pherones/blood-bank-management.git/utils"
)

func AuthMiddleware(expectedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.SendError(c, fiber.StatusUnauthorized, "Authorization header is missing", nil)
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return utils.SendError(c, fiber.StatusUnauthorized, "Invalid token format", nil)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		accessSecret := os.Getenv("JWT_ACCESS_SECRET")
		claims, err := utils.VerifyToken(tokenString, accessSecret)
		if err != nil {
			return utils.SendError(c, fiber.StatusUnauthorized, "Invalid or expired access token", err)
		}

		userRole := claims["role"].(string)

		roleMatch := false
		for _, role := range expectedRoles {
			if userRole == role {
				roleMatch = true
				break
			}
		}

		if !roleMatch {
			return utils.SendError(c, fiber.StatusForbidden, "You do not have the required permissions", nil)
		}

		c.Locals("userID", claims["user_id"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}
