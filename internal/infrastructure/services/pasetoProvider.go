package services

import (
	"time"

	"github.com/google/uuid"
	. "github.com/kerimcetinbas/go_ddd_ca/application/common/interfaces/auth"
	. "github.com/kerimcetinbas/go_ddd_ca/domain/auth"
	. "github.com/kerimcetinbas/go_ddd_ca/domain/user"
	"github.com/o1egl/paseto"
)

type pasetoTokenProvider struct {
	settigns PasetoTokenSettings
}

func UsePasetoTokenProvider(settings PasetoTokenSettings) ITokenProvider {
	return &pasetoTokenProvider{
		settigns: settings,
	}
}

// GenerateToken implements interfaces_authentication.IPasetoTokenGenerator.
func (p *pasetoTokenProvider) GenerateAccessToken(user User) (string, error) {

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

func (p *pasetoTokenProvider) GenerateRefreshToken(user User) (string, error) {
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
