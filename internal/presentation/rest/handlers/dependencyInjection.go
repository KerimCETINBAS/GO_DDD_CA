package handlers

import (
	auth_handler "gihub.com/kerimcetinbas/go_ddd_ca/presentation/rest/handlers/auth"
	user_handler "gihub.com/kerimcetinbas/go_ddd_ca/presentation/rest/handlers/users"
	"go.uber.org/dig"
)

func UseRouters(c *dig.Container) {
	c.Invoke(auth_handler.AuthRouter)
	c.Invoke(user_handler.UserRouter)
}

func UseHandlers(c *dig.Container) {
	c.Provide(auth_handler.AuthHandler)
	c.Provide(user_handler.UserHandler)
}
