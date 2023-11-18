package main

import (
	"github.com/gofiber/fiber/v2"
	application "github.com/kerimcetinbas/go_ddd_ca/application"
	domain "github.com/kerimcetinbas/go_ddd_ca/domain"
	"github.com/kerimcetinbas/go_ddd_ca/infrastructure"
	presentation "github.com/kerimcetinbas/go_ddd_ca/presentation"
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
