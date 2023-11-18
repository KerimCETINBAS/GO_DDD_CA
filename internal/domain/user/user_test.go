package user_domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	user_valueobject "github.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
	"github.com/stretchr/testify/assert"
)

func TestUserAggregate(t *testing.T) {

	t.Run("Must be able to create a new user with a unique ID", func(t *testing.T) {

		user := NewUser(
			"jdoe@example.com",
			"jhon doe",
			[]byte("1234"),
			make([]user_valueobject.FriendIdValueObject, 0))

		_, ok := uuid.Parse(user.Id().Value().String())

		assert.Nil(t, ok)
		assert.Equal(t, "jdoe@example.com", user.Email())
		assert.Equal(t, "jdoe@example.com", user.email)
		assert.Equal(t, "jhon doe", user.UserName())
		assert.Equal(t, "jhon doe", user.userName)
		assert.True(t, time.Now().After(user.CreatedAt()))
		assert.True(t, time.Now().After(user.UpdatedAt()))
		assert.Equal(t, []byte("1234"), user.Password())
		assert.Equal(t, []byte("1234"), user.password)
		assert.Equal(t, make([]user_valueobject.FriendIdValueObject, 0), user.FriendIds())
		assert.Equal(t, make([]user_valueobject.FriendIdValueObject, 0), user.friendIds)
	})
}
