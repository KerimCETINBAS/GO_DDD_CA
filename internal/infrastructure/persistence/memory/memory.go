package memory

import (
	"go.uber.org/dig"
)

func UseMemoryProviders(c *dig.Container) {
	c.Provide(userMemRepositoryProvider)
}
