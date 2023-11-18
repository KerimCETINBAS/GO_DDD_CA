package auth

import (
	user_domain "github.com/kerimcetinbas/go_ddd_ca/domain/user"
)

type ITokenProvider interface {
	GenerateAccessToken(user_domain.User) (string, error)
}
