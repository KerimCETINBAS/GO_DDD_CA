package auth_test

import (
	"testing"
	"time"

	"github.com/joho/godotenv"
	. "github.com/kerimcetinbas/go_ddd_ca/domain/auth"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

type testCase struct {
	name     string
	input    []string
	expected interface{}
}

type testCases []testCase

func TestPasetoSettings(t *testing.T) {
	t.Run("Testing Paseto Provider", func(t *testing.T) {
		asserts := assert.New(t)
		godotenv.Load("../../../.env.development")

		c := dig.New()

		c.Provide(PasetoTokenSettingsProvider)

		c.Invoke(func(settings PasetoTokenSettings) {
			asserts.Equal(settings.Access.Audience, "ZEHIR_ZIKKIM")
			asserts.Equal(settings.Access.SymmetricKey, []byte("super-secret-keysdf2353234234232"))
			asserts.Equal(settings.Access.Issuer, "ZEHIR_ZIKKIM")
			asserts.Equal(settings.Access.ExpiresAfter, time.Duration(20*time.Minute))
			asserts.Equal(settings.Access.ContextKey, "User")
		})
	})
}
