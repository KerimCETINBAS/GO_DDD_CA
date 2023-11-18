package persistence

import (
	user_entity "gihub.com/kerimcetinbas/go_ddd_ca/domain/user/entities"
	user_valueobject "gihub.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
)

type IFriendRepository interface {
	Create(user_entity.Friend) error
	List() ([]user_entity.Friend, error)
	Delete(user_valueobject.FriendIdValueObject) error
}
