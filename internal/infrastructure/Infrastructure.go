package infrastructure

import (
	"gihub.com/kerimcetinbas/go_ddd_ca/infrastructure/persistence/memory"
	"gihub.com/kerimcetinbas/go_ddd_ca/infrastructure/services"
	"go.uber.org/dig"
)

type ProviderFactory = func(*dig.Container)

func AddInfrastructure(container *dig.Container) {

	services.UseEnvProvider(container)

	container.Provide(services.UsePasetoProvider)
	container.Provide(services.UseDateTimeProvider)

	container.Invoke(memory.UseMemoryProviders)
}
