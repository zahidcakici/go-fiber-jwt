package middleware

import (
	"go-fiber-jwt-example/config"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

func JWTProtected(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Get("JWT_SECRET"))},
		ContextKey: "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Return status 401 and failed authentication error.
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		},
	})(c)
}
