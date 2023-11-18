package auth

import (
	user_domain "gihub.com/kerimcetinbas/go_ddd_ca/domain/user"
)

type IPasetoTokenGenerator interface {
	GenerateAccessToken(user_domain.User) (string, error)
}
