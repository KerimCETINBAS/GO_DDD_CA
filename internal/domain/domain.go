package domain

import (
	. "gihub.com/kerimcetinbas/go_ddd_ca/domain/auth"
	"go.uber.org/dig"
)

func AddDomain(container *dig.Container) {
	container.Provide(PasetoTokenSettingsProvider)
}
