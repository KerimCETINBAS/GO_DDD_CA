package domain_errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainErrors(t *testing.T) {

	t.Run("Must be throw correct type of errors", func(t *testing.T) {

		expect := assert.New(t)

		expect.Equal(ThrowError().ERR_ConflictException, ERR_ConflictException)
		expect.Equal(ThrowError().ERR_InsufficientAuthorityException, ERR_InsufficientAuthorityException)
		expect.Equal(ThrowError().ERR_MalformedEntityException, ERR_MalformedEntityException)
		expect.Equal(ThrowError().ERR_NoAnonymousAllowedException, ERR_NoAnonymousAllowedException)
		expect.Equal(ThrowError().ERR_PasswordDoesNotMatchException, ERR_PasswordDoesNotMatchException)
		expect.Equal(ThrowError().ERR_UserAlreadyExistException, ERR_UserAlreadyExistException)
		expect.Equal(ThrowError().ERR_UserDoesNotExistException, ERR_UserDoesNotExistException)
		expect.Equal(ThrowError().ERR_UserEmailAlreadyExistException, ERR_UserEmailAlreadyExistException)
	})
}
