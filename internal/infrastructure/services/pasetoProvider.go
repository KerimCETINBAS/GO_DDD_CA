package services

import (
	"time"

	. "gihub.com/kerimcetinbas/go_ddd_ca/application/common/interfaces/auth"
	. "gihub.com/kerimcetinbas/go_ddd_ca/domain/auth"
	. "gihub.com/kerimcetinbas/go_ddd_ca/domain/user"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type pasetoProvider struct {
	settigns PasetoTokenSettings
}

func UsePasetoProvider(settings PasetoTokenSettings) IPasetoTokenGenerator {
	return &pasetoProvider{
		settigns: settings,
	}
}

// GenerateToken implements interfaces_authentication.IPasetoTokenGenerator.
func (p *pasetoProvider) GenerateAccessToken(user User) (string, error) {

	payload := PasetoPayload{
		JSONToken: paseto.JSONToken{
			Jti:        uuid.NewString(),
			Audience:   p.settigns.Access.Audience,
			Expiration: time.Now().Add(p.settigns.Access.ExpiresAfter),
			Issuer:     p.settigns.Access.Issuer,
			IssuedAt:   time.Now(),
			Subject:    user.Email(),
		},
	}

	return paseto.NewV2().Encrypt(p.settigns.Access.SymmetricKey, payload, nil)
}

func (p *pasetoProvider) GenerateRefreshToken(user User) (string, error) {
	payload := PasetoPayload{
		JSONToken: paseto.JSONToken{
			Jti:        uuid.NewString(),
			Audience:   p.settigns.Access.Audience,
			Expiration: time.Now().Add(time.Second * 10),
			Issuer:     p.settigns.Access.Issuer,
			IssuedAt:   time.Now(),
			Subject:    user.Email(),
		},
	}

	return paseto.NewV2().Encrypt(p.settigns.Refresh.SymmetricKey, payload, nil)
}
