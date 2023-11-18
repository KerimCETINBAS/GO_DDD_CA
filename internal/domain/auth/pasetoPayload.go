package auth

import (
	"github.com/o1egl/paseto"
)

type PasetoPayload struct {
	paseto.JSONToken
}
