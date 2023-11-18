package login_query

import (
	"context"
	"errors"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/kerimcetinbas/go_ddd_ca/application/common/services"
	domain_errors "github.com/kerimcetinbas/go_ddd_ca/domain/common/Errors"
	user_domain "github.com/kerimcetinbas/go_ddd_ca/domain/user"
	user_valueobject "github.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
	. "github.com/kerimcetinbas/go_ddd_ca/infrastructure/services"
	"github.com/kerimcetinbas/go_ddd_ca/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func TestRegisterCommandHandler(t *testing.T) {
	godotenv.Load("../../.env.development")

	c := dig.New()

	c.Provide(UseDateTimeProvider)
	mockTokenProvider := new(mocks.MockTokenProvider)
	mockRepo := new(mocks.MockUserRepository)

	c.Provide(func(dateProvider IDateTimeProvider) *LoginUserQueryHandler {
		return &LoginUserQueryHandler{
			dateTimeProvider: dateProvider,
			tokenGenerator:   mockTokenProvider,
			userRepository:   mockRepo,
		}
	})
	t.Run("Should have user repository", func(t *testing.T) {

		c.Invoke(func(handler *LoginUserQueryHandler) {
			assert.NotNil(t, handler.userRepository)
		})
	})

	t.Run("Should login user with correct credentials", func(t *testing.T) {
		c.Invoke(func(handler LoginUserQueryHandler) {
			q := &LoginUserQuery{
				Email:    "jdoe@example.com",
				Password: []byte("1234"),
			}

			u := user_domain.NewUser(
				q.Email,
				"jhond",
				q.Password,
				[]user_valueobject.FriendIdValueObject{},
			)
			mockRepo.On("GetByEmail", q.Email).Return(u, nil)
			mockTokenProvider.On("GenerateAccessToken", u).Return("token", nil)

			user, err := handler.Handle(context.Background(), q)

			assert.Equal(t, u.Id().Value(), user.Id)
			assert.Equal(t, user.AccessToken, "token")
			assert.Equal(t, u.Email(), user.Email)
			assert.Equal(t, u.UserName(), user.UserName)
			assert.NotNil(t, user)
			assert.Nil(t, err)

		})
	})

	t.Run("Should throw error with incorrect credentials", func(t *testing.T) {

		c.Invoke(func(handler *LoginUserQueryHandler) {

			q := &LoginUserQuery{
				Email:    "jdoe@example.com",
				Password: []byte("1234"),
			}

			u := user_domain.User{}

			mockRepo.On("GetByEmail", q.Email).Return(u, domain_errors.ThrowError().ERR_UserDoesNotExistException)

			user, err := handler.Handle(context.Background(), q)

			assert.Equal(t, user, &LoginUserQueryResponse{})
			assert.NotNil(t, err)
			assert.True(t, errors.Is(err, domain_errors.ERR_UserDoesNotExistException))

		})
	})
}
