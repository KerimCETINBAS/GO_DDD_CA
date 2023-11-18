package persistence

import (
	user_entity "github.com/kerimcetinbas/go_ddd_ca/domain/user/entities"
	user_valueobject "github.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
)

type IFriendRepository interface {
	Create(user_entity.Friend) error
	List() ([]user_entity.Friend, error)
	Delete(user_valueobject.FriendIdValueObject) error
}
