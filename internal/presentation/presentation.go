package presentation

import (
	"gihub.com/kerimcetinbas/go_ddd_ca/presentation/rest"
	"gihub.com/kerimcetinbas/go_ddd_ca/presentation/rest/handlers"
	"gihub.com/kerimcetinbas/go_ddd_ca/presentation/rest/middlewares"
	"go.uber.org/dig"
)

func AddPresentation(c *dig.Container) {
	c.Provide(rest.Rest)
	c.Provide(middlewares.AuthMiddlewareProvider)
	handlers.UseHandlers(c)
	handlers.UseRouters(c)

}
