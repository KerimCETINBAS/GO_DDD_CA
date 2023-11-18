package middlewares

import (
	"encoding/json"
	"errors"
	"time"

	. "gihub.com/kerimcetinbas/go_ddd_ca/domain/auth"
	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/o1egl/paseto"
)

func pasetoErrorHandler(c *fiber.Ctx, err error) error {

	if errors.Is(err, pasetoware.ErrMissingToken) {
		return c.Next()
	} else if errors.Is(err, pasetoware.ErrIncorrectTokenPrefix) {
		return fiber.ErrUnprocessableEntity
	} else if errors.Is(err, pasetoware.ErrDataUnmarshal) {
		return fiber.ErrUnprocessableEntity
	} else if errors.Is(err, paseto.ErrInvalidTokenAuth) {
		return fiber.ErrUnauthorized
	} else if errors.Is(err, paseto.ErrInvalidSignature) {
		return fiber.ErrUnprocessableEntity
	}

	return fiber.ErrInternalServerError
}

func pasetoTokenValidate(decrypted []byte) (interface{}, error) {
	payload := PasetoPayload{}
	err := json.Unmarshal(decrypted, &payload)

	if payload.Expiration.Before(time.Now()) {
		return payload, fiber.ErrUnauthorized
	}
	return payload, err
}

func AccessTokenMiddleware(pasetoTokenSettings PasetoTokenSettings) func(*fiber.Ctx) error {
	return pasetoware.New(pasetoware.Config{
		SymmetricKey: pasetoTokenSettings.Access.SymmetricKey,
		TokenLookup:  pasetoTokenSettings.Access.TokenLookup,
		ContextKey:   pasetoTokenSettings.Access.ContextKey,
		TokenPrefix:  pasetoTokenSettings.Access.TokenPrefix,
		ErrorHandler: pasetoErrorHandler,
		Validate:     pasetoTokenValidate,
	})
}

func RefreshTokenMiddleware(pasetoTokenSettings PasetoTokenSettings) func(*fiber.Ctx) error {
	return pasetoware.New(pasetoware.Config{
		SymmetricKey: pasetoTokenSettings.Refresh.SymmetricKey,
		TokenLookup:  pasetoTokenSettings.Refresh.TokenLookup,
		ContextKey:   pasetoTokenSettings.Refresh.ContextKey,
		TokenPrefix:  pasetoTokenSettings.Refresh.TokenPrefix,
		ErrorHandler: pasetoErrorHandler,
		Validate:     pasetoTokenValidate,
	})
}
