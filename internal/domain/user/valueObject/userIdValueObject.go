package user_valueobject

import (
	"github.com/google/uuid"
	"github.com/kerimcetinbas/go_ddd_ca/domain/common/models"
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
