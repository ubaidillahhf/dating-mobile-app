package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/dating-service/app/infra/config"
	"github.com/ubaidillahhf/dating-service/app/infra/presenter"
)

func ValidateToken(c *fiber.Ctx) error {
	secret := config.GetEnv("ACCESS_TOKEN_SECRET")
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		res := presenter.Unauthorize("error: please provide valid token!", nil)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	token := strings.Split(authHeader, " ")[1]

	isAuthorized, iaErr := IsAuthorized(token, secret)
	if iaErr != nil {
		res := presenter.Unauthorize("error: invalid or expired token!", nil)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	if isAuthorized {
		userId, uiErr := ExtractIDFromToken(token, secret)
		if uiErr != nil {
			res := presenter.Unauthorize("error: unknows user!", nil)
			return c.Status(fiber.StatusUnauthorized).JSON(res)
		}

		c.Locals("myId", userId)
	}

	return c.Next()
}
