package presentation

import (
	"github.com/kerimcetinbas/go_ddd_ca/presentation/rest"
	"github.com/kerimcetinbas/go_ddd_ca/presentation/rest/handlers"
	"github.com/kerimcetinbas/go_ddd_ca/presentation/rest/middlewares"
	"go.uber.org/dig"
)

func AddPresentation(c *dig.Container) {
	c.Provide(rest.Rest)
	c.Provide(middlewares.AuthMiddlewareProvider)
	handlers.UseHandlers(c)
	handlers.UseRouters(c)

}
