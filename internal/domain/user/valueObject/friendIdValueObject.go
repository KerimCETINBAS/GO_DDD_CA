package user_valueobject

import (
	"github.com/google/uuid"
	"github.com/kerimcetinbas/go_ddd_ca/domain/common/models"
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
