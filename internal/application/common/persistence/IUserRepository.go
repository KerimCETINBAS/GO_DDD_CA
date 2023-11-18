package persistence

import (
	user_domain "github.com/kerimcetinbas/go_ddd_ca/domain/user"
	user_valueobject "github.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
)

type IUserRepository interface {
	Create(user_domain.User) error
	GetById(user_valueobject.UserIdValueObject) (user_domain.User, error)
	GetByEmail(email string) (user_domain.User, error)
	GetAll() ([]user_domain.User, error)
}
