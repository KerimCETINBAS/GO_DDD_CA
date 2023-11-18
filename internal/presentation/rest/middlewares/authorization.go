package middlewares

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/kerimcetinbas/go_ddd_ca/domain/auth"
)

type IAuthMiddleware interface {
	Validate() fiber.Handler
	ValidateByRole() fiber.Handler
	ValidateByPerms() fiber.Handler
}

type authMiddleware struct {
	settings PasetoTokenSettings
}

func AuthMiddlewareProvider(pasetoSettings PasetoTokenSettings) IAuthMiddleware {
	return &authMiddleware{
		settings: pasetoSettings,
	}
}

func (m *authMiddleware) Validate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := c.Locals(m.settings.Access.ContextKey).(PasetoPayload)

		if &payload.Subject == nil || payload.Subject == "" {
			return fiber.ErrUnauthorized
		}

		return c.Next()
	}
}

func (m *authMiddleware) ValidateByRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func (m *authMiddleware) ValidateByPerms() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
