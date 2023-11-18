package memory

import (
	"sync"

	"gihub.com/kerimcetinbas/go_ddd_ca/application/common/persistence"
	domain_errors "gihub.com/kerimcetinbas/go_ddd_ca/domain/common/Errors"
	user_domain "gihub.com/kerimcetinbas/go_ddd_ca/domain/user"
	user_valueobject "gihub.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
)

type userRepository struct {
	users *[]user_domain.User
	sync.Mutex
}

func userMemRepositoryProvider() persistence.IUserRepository {
	usrs := make([]user_domain.User, 0)
	return &userRepository{
		users: &usrs,
	}
}

func Find[T interface{}](slice []T, prediction func(el T, index int, slice []T) bool) (*T, domain_errors.IError) {

	for in, el := range slice {
		if prediction(el, in, slice) {
			return &el, nil
		}
	}
	return new(T), domain_errors.ThrowError().ERR_UserDoesNotExistException
}

// Create implements persistence.IUserRepository.
func (r *userRepository) Create(user user_domain.User) error {

	if _, err := Find[user_domain.User](*r.users, func(el user_domain.User, _ int, _ []user_domain.User) bool {
		return el.Email() == user.Email()
	}); err != nil {
		r.Lock()
		*r.users = append(*r.users, user)
		r.Unlock()

		return nil
	}

	return domain_errors.ThrowError().ERR_UserAlreadyExistException
}

// GetAll implements persistence.IUserRepository.
func (r *userRepository) GetAll() ([]user_domain.User, error) {
	return *r.users, nil
}

// GetByEmail implements persistence.IUserRepository.
func (r *userRepository) GetByEmail(email string) (user_domain.User, error) {

	user, err := Find[user_domain.User](*r.users, func(el user_domain.User, _ int, _ []user_domain.User) bool {
		return el.Email() == email
	})

	return *user, err
}

// GetById implements persistence.IUserRepository.
func (r *userRepository) GetById(id user_valueobject.UserIdValueObject) (user_domain.User, error) {
	user, err := Find[user_domain.User](*r.users, func(el user_domain.User, _ int, _ []user_domain.User) bool {
		return el.Id().ValueObjet.Compare(id)
	})

	if err != nil {
		return *user, domain_errors.ThrowError().ERR_UserDoesNotExistException
	}

	return *user, nil
}
