package user_entity

import (
	"github.com/kerimcetinbas/go_ddd_ca/domain/common/models"
	user_valueobject "github.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
)

type Friend struct {
	models.Entity[user_valueobject.FriendIdValueObject]
	UserName string
}

func NewFriend(userName string) *Friend {
	return &Friend{
		Entity: models.Entity[user_valueobject.FriendIdValueObject]{
			Id: user_valueobject.NewFriendId(),
		},
		UserName: userName,
	}
}
