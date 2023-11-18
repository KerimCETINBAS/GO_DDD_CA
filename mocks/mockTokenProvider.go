package mock

import (
	user_domain "gihub.com/kerimcetinbas/go_ddd_ca/domain/user"
	"github.com/stretchr/testify/mock"
)

type MockTokenProvider struct {
	mock.Mock
}

func (m *MockTokenProvider) GenerateAccessToken(u user_domain.User) (string, error) {
	args := m.Called(u)
	return args.String(0), args.Error(1)
}
