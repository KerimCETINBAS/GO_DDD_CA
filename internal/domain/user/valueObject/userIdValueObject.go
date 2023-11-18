package user_valueobject

import (
	"gihub.com/kerimcetinbas/go_ddd_ca/domain/common/models"
	"github.com/google/uuid"
)

type UserIdValueObject struct {
	models.ValueObjet[UserIdValueObject]
	value uuid.UUID
}

func (v *UserIdValueObject) Value() uuid.UUID {
	return v.value
}

func NewUserId(value uuid.UUID) UserIdValueObject {
	return UserIdValueObject{
		ValueObjet: models.ValueObjet[UserIdValueObject]{},
		value:      value,
	}
}
