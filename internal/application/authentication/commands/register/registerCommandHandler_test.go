package register_command

import (
	"context"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kerimcetinbas/go_ddd_ca/application/common/persistence"
	. "github.com/kerimcetinbas/go_ddd_ca/application/common/services"
	user_domain "github.com/kerimcetinbas/go_ddd_ca/domain/user"
	user_valueobject "github.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
	"github.com/kerimcetinbas/go_ddd_ca/infrastructure/services"
	mocks "github.com/kerimcetinbas/go_ddd_ca/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/dig"
)

func TestRegisterCommandHandler(t *testing.T) {
	err := godotenv.Load("../../.env.development")

	fmt.Println(err)
	c := dig.New()

	c.Provide(services.UseDateTimeProvider)
	// user := user_domain.NewUser(
	// 	"jdoe@example.com",
	// 	"jhond",
	// 	[]byte("1234"),
	// 	[]user_valueobject.FriendIdValueObject{
	// 		user_valueobject.NewFriendId(),
	// 	},
	// )

	mockRepo := new(mocks.MockUserRepository)
	c.Provide(func() persistence.IUserRepository {
		return mockRepo
	})
	c.Provide(func(dateProvider IDateTimeProvider) RegisterUserCommandHandler {
		return RegisterUserCommandHandler{
			DateTimeProvider: dateProvider,
			userRespository:  mockRepo,
		}
	})
	t.Run("Should have user repository", func(t *testing.T) {
		c.Invoke(func(handler *RegisterUserCommandHandler) {
			assert.NotNil(t, handler.userRespository)
		})
	})

	t.Run("Should register user", func(t *testing.T) {
		c.Invoke(func(handler RegisterUserCommandHandler) {

			r := &RegisterUserCommand{
				UserName: "jhond",
				Email:    "jdoe@example.com",
				Password: []byte("1234"),
			}
			u := user_domain.NewUser(
				r.Email,
				r.UserName,
				r.Password,
				make([]user_valueobject.FriendIdValueObject, 0),
			)

			mockRepo.On("Create", mock.MatchedBy(func(s user_domain.User) bool {
				return s.UserName() == u.UserName()
			})).Return(nil)
			_, err := handler.Handle(context.Background(), r)

			assert.Nil(t, err)

		})
	})
}
