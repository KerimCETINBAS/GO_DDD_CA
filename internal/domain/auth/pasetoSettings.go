package auth

import (
	"os"
	"strconv"
	"time"
)

type TokenSetting struct {
	SymmetricKey []byte
	Audience     string
	Issuer       string
	ExpiresAfter time.Duration
	NotBefore    time.Duration
	TokenLookup  [2]string
	ContextKey   string
	TokenPrefix  string
}

type PasetoTokenSettings struct {
	Access  TokenSetting
	Refresh TokenSetting
}

func PasetoTokenSettingsProvider() PasetoTokenSettings {

	minAccess, _ := strconv.ParseUint(os.Getenv("PASETO_ACCESS_EXPIRES_TIME_IN_MINUTES"), 10, 64)
	minRefresh, _ := strconv.ParseUint(os.Getenv("PASETO_REFRESH_EXPIRES_TIME_IN_MINUTES"), 10, 64)
	return PasetoTokenSettings{
		Access: TokenSetting{
			SymmetricKey: []byte(os.Getenv("PASETO_ACCESS_SYMMETRIC_KEY")),
			Audience:     os.Getenv("PASETO_AUDIENCE"),
			Issuer:       os.Getenv("PASETO_ISSUER"),
			ExpiresAfter: time.Duration(time.Minute * time.Duration(minAccess)),
			ContextKey:   "User",
			TokenLookup:  [2]string{"header", "Authorization"},
			TokenPrefix:  "Bearer",
		},

		Refresh: TokenSetting{
			SymmetricKey: []byte(os.Getenv("PASETO_REFRESH_SYMMETRIC_KEY")),
			Audience:     os.Getenv("PASETO_AUDIENCE"),
			Issuer:       os.Getenv("PASETO_ISSUER"),
			ExpiresAfter: time.Duration(time.Minute * time.Duration(minRefresh)),
			ContextKey:   "User",
			TokenLookup:  [2]string{"header", "Authorization"},
			TokenPrefix:  "Bearer",
		},
	}
}
