package rest

import (
	. "gihub.com/kerimcetinbas/go_ddd_ca/domain/auth"
	"gihub.com/kerimcetinbas/go_ddd_ca/presentation/rest/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Rest(pasetoSettings PasetoTokenSettings) *fiber.App {
	app := fiber.New()

	app.Use(idempotency.New())
	app.Use(helmet.New())
	app.Use(requestid.New())
	app.Use(middlewares.AccessTokenMiddleware(pasetoSettings))
	return app
}
