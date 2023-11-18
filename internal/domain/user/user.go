package user_domain

import (
	"github.com/google/uuid"
	"github.com/kerimcetinbas/go_ddd_ca/domain/common/models"
	user_valueobject "github.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
)

type User struct {
	models.AggregateRoot[user_valueobject.UserIdValueObject]
	email     string
	userName  string
	password  []byte
	friendIds []user_valueobject.FriendIdValueObject
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) UserName() string {
	return u.userName
}

func (u *User) SetUserName(userName string) {
	u.userName = userName
}

func (u *User) Password() []byte {
	return u.password
}

func (u *User) FriendIds() []user_valueobject.FriendIdValueObject {
	return u.friendIds
}

func (u *User) AddFriend(friendId user_valueobject.FriendIdValueObject) {
	u.friendIds = append(u.friendIds, friendId)
}

func (u *User) RemoveFriend(friendId user_valueobject.FriendIdValueObject) {

	nFriends := make([]user_valueobject.FriendIdValueObject, 0)

	for _, id := range u.friendIds {

		if !id.Compare(friendId) {
			nFriends = append(nFriends, id)
		}
	}
}

func NewUser(
	email string,
	userName string,
	password []byte,
	friendIds []user_valueobject.FriendIdValueObject) User {
	return User{
		AggregateRoot: models.
			NewAggregateRoot[user_valueobject.UserIdValueObject](
			user_valueobject.NewUserId(uuid.New())),
		email:     email,
		userName:  userName,
		password:  password,
		friendIds: friendIds,
	}
}
