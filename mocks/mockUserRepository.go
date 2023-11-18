package mocks

import (
	"github.com/kerimcetinbas/go_ddd_ca/application/common/persistence"
	. "github.com/kerimcetinbas/go_ddd_ca/domain/user"
	. "github.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

// Create implements persistence.IUserRepository.
func (m *MockUserRepository) Create(user User) error {

	args := m.Called(user)

	return args.Error(0)
}

// GetAll implements persistence.IUserRepository.
func (m *MockUserRepository) GetAll() ([]User, error) {
	args := m.Called()
	return args.Get(0).([]User), args.Error(1)
}

// GetByEmail implements persistence.IUserRepository.
func (m *MockUserRepository) GetByEmail(email string) (User, error) {
	args := m.Called(email)

	return args.Get(0).(User), args.Error(1)
}

// GetById implements persistence.IUserRepository.
func (*MockUserRepository) GetById(UserIdValueObject) (User, error) {
	panic("unimplemented")
}

func MockUserRepositoryConstructor() persistence.IUserRepository {
	return &MockUserRepository{}
}
