package application

import (
	"go.uber.org/dig"
)

func AddApplication(container *dig.Container) {
	InjectCommands(container)
	InjectQueries(container)

}
