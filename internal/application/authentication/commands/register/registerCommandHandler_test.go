package register_command

import (
	"context"
	"fmt"
	"testing"

	"gihub.com/kerimcetinbas/go_ddd_ca/application/common/persistence"
	. "gihub.com/kerimcetinbas/go_ddd_ca/application/common/services"
	"gihub.com/kerimcetinbas/go_ddd_ca/domain/auth"
	"gihub.com/kerimcetinbas/go_ddd_ca/infrastructure/services"
	"github.com/joho/godotenv"
	mock "github.com/kerimcetinbas/go_ddd_ca/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func TestRegisterCommandHandler(t *testing.T) {
	err := godotenv.Load("../../.env.development")

	fmt.Println(err)
	c := dig.New()

	c.Provide(auth.PasetoTokenSettingsProvider)
	c.Provide(services.UseDateTimeProvider)
	c.Provide(services.UsePasetoProvider)
	// user := user_domain.NewUser(
	// 	"jdoe@example.com",
	// 	"jhond",
	// 	[]byte("1234"),
	// 	[]user_valueobject.FriendIdValueObject{
	// 		user_valueobject.NewFriendId(),
	// 	},
	// )

	mockRepo := new(mock.MockUserRepository)
	c.Provide(func() persistence.IUserRepository {
		return mockRepo
	})
	c.Provide(func(dateProvider IDateTimeProvider, userRepo persistence.IUserRepository) *RegisterUserCommandHandler {
		return &RegisterUserCommandHandler{
			DateTimeProvider: dateProvider,
			userRespository:  userRepo,
		}
	})
	t.Run("Should have user repository", func(t *testing.T) {
		c.Invoke(func(handler *RegisterUserCommandHandler) {
			assert.NotNil(t, handler.userRespository)
		})
	})

	t.Run("Should register user", func(t *testing.T) {
		c.Invoke(func(handler *RegisterUserCommandHandler) {
			_, err := handler.Handle(context.Background(), &RegisterUserCommand{
				UserName: "jhond",
				Email:    "jdoe@example.com",
				Password: []byte("1234"),
			})

			assert.Nil(t, err)

		})
	})
}
