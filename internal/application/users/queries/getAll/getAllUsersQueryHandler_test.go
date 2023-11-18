package users_getall_query

import (
	"context"
	"testing"

	"gihub.com/kerimcetinbas/go_ddd_ca/application/common/persistence"
	user_domain "gihub.com/kerimcetinbas/go_ddd_ca/domain/user"
	user_valueobject "gihub.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"

	mock "github.com/kerimcetinbas/go_ddd_ca/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func TestGetAllUsersQueryHandler(t *testing.T) {

	c := dig.New()

	user := user_domain.NewUser(
		"jdoe@example.com",
		"jhond",
		[]byte("1234"),
		[]user_valueobject.FriendIdValueObject{
			user_valueobject.NewFriendId(),
		},
	)

	mockRepo := new(mock.MockUserRepository)
	c.Provide(func() persistence.IUserRepository {
		return mockRepo
	})
	t.Run("Should have user repository", func(t *testing.T) {
		c.Invoke(func(handler *GetAllUserQueryHandler) {
			assert.NotNil(t, handler.userRepository)
		})
	})

	t.Run("it should get empty slice", func(t *testing.T) {
		c.Invoke(func(handler *GetAllUserQueryHandler, mock persistence.IUserRepository) {
			mockRepo.AssertNumberOfCalls(t, "GetAll", 0)
			expect, err := handler.Handle(context.TODO(), &GetAllUsersQuery{})
			mockRepo.On("GetAll").Return([]user_domain.User{}, nil)
			assert.Nil(t, err)
			assert.Equal(t, expect, &GetAllUsersQueryResponse{})

		})
	})

	t.Run("Should get all new users", func(t *testing.T) {
		c.Invoke(func(handler *GetAllUserQueryHandler, mock persistence.IUserRepository) {
			expect, err := handler.Handle(context.Background(), &GetAllUsersQuery{})
			mockRepo.On("GetAll").Return([]user_domain.User{user}, nil)
			assert.Nil(t, err)
			assert.Equal(t, expect, &GetAllUsersQueryResponse{
				{
					Id:       user.Id().Value().String(),
					UserName: user.UserName(),
					Email:    user.Email(),
				},
			})
		})
	})

}
