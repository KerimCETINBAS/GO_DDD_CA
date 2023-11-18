package main

import (
	application "gihub.com/kerimcetinbas/go_ddd_ca/application"
	domain "gihub.com/kerimcetinbas/go_ddd_ca/domain"
	"gihub.com/kerimcetinbas/go_ddd_ca/infrastructure"
	presentation "gihub.com/kerimcetinbas/go_ddd_ca/presentation"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func main() {
	Bootstrap(dig.New())
}

func Bootstrap(c *dig.Container) {
	c.Provide(func() *dig.Container {
		return c
	})

	domain.AddDomain(c)
	infrastructure.AddInfrastructure(c)
	application.AddApplication(c)
	presentation.AddPresentation(c)

	c.Invoke(func(app *fiber.App) {
		app.Listen(":8080")
	})

}
