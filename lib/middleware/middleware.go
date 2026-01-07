package middleware

import (
	"news-web/config"
	"news-web/internal/adapter/handler/response"
	"news-web/lib/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Middleware interface {
	CheckToken() fiber.Handler
}

type authMiddleware struct {
	authJwt auth.Jwt
}

// CheckToken implements Middleware.
func (m *authMiddleware) CheckToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHandler := c.Get("Authorization")

		var errorResponse response.ErrorResponseDefault
		if authHandler == "" {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Authorization header missing"
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header missing",
			})
		}

		tokenString := strings.Split(authHandler, "Bearer ")[1]
		claims, err := m.authJwt.VerifyAccessToken(tokenString)
		if err != nil {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Invalid or expired token"
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		c.Locals("user", claims)
		return c.Next()
	}
}

func NewMiddleware(cfg *config.Config) Middleware {
	opt := new(authMiddleware)
	opt.authJwt = auth.NewJwt(cfg)
	return opt
}
