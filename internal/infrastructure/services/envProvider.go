package services

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kerimcetinbas/go_ddd_ca/application/common/services"
	"go.uber.org/dig"
)

type envProvider struct {
}

func (e *envProvider) Get(key string) string {
	return os.Getenv(key)
}

func UseEnvProvider(c *dig.Container) {

	d, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprint(d, "/", ".env", ".", os.Getenv("GO_ENV")))

	if err != nil {
		panic(err)
	}

	c.Provide(func() services.IEnvProvider {
		return &envProvider{}
	})
}
