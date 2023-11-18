package domain

import (
	. "github.com/kerimcetinbas/go_ddd_ca/domain/auth"
	"go.uber.org/dig"
)

func AddDomain(container *dig.Container) {
	container.Provide(PasetoTokenSettingsProvider)
}
