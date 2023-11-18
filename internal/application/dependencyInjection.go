package application

import (
	register_command "gihub.com/kerimcetinbas/go_ddd_ca/application/authentication/commands/register"
	login_query "gihub.com/kerimcetinbas/go_ddd_ca/application/authentication/queries/login"
	users_getall_query "gihub.com/kerimcetinbas/go_ddd_ca/application/users/queries/getAll"
	"go.uber.org/dig"
)

func InjectCommands(c *dig.Container) {

	commandProviders := [][]interface{}{
		{
			register_command.Provider,
			register_command.Invoke,
		},
	}

	for _, p := range commandProviders {
		c.Provide(p[0])
		c.Invoke(p[1])
	}

}

func InjectQueries(c *dig.Container) {
	queryProviders := [][]interface{}{
		{
			login_query.Provider,
			login_query.Invoke,
		},
		{
			users_getall_query.Provider,
			users_getall_query.Invoke,
		},
	}

	for _, p := range queryProviders {
		c.Provide(p[0])
		c.Invoke(p[1])
	}
}
