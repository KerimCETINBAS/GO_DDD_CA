package user_valueobject

import (
	"gihub.com/kerimcetinbas/go_ddd_ca/domain/common/models"
	"github.com/google/uuid"
)

type FriendIdValueObject struct {
	models.ValueObjet[FriendIdValueObject]
	value uuid.UUID
}

func (v *FriendIdValueObject) Value() uuid.UUID {
	return v.value
}

func NewFriendId() FriendIdValueObject {
	return FriendIdValueObject{
		ValueObjet: models.ValueObjet[FriendIdValueObject]{},
	}
}
