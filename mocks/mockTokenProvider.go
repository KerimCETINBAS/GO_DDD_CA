package mocks

import (
	"time"

	. "github.com/kerimcetinbas/go_ddd_ca/domain/auth"
	. "github.com/kerimcetinbas/go_ddd_ca/domain/user"
	"github.com/stretchr/testify/mock"
)

type MockTokenProvider struct {
	mock.Mock
}

var MockTokenSettings = PasetoTokenSettings{
	Access: TokenSetting{
		SymmetricKey: []byte("super-secret-key-1234-key-1-1234"),
		Audience:     "dd_test",
		Issuer:       "dd_test",
		ExpiresAfter: time.Duration(time.Minute * 10),
		ContextKey:   "User",
		TokenLookup:  [2]string{"header", "Authorization"},
		TokenPrefix:  "Bearer",
	},
}

func (m *MockTokenProvider) GenerateAccessToken(u User) (string, error) {
	args := m.Called(u)
	return args.String(0), args.Error(1)
}
