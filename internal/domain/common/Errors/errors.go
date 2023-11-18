package domain_errors

import "errors"

// common

type IError interface {
	Error() string
}

var (
	ERR_ConflictException              = errors.New("Conflict excetpion")
	ERR_ValidationExeption             = errors.New("Validation exception")
	ERR_MalformedEntityException       = errors.New("Malformed entity exception")
	ERR_UserDoesNotExistException      = errors.New("User doest not exist.")
	ERR_UserEmailAlreadyExistException = errors.New("User email already exists.")
	ERR_UserAlreadyExistException      = errors.New("User already exists.")
	ERR_PasswordDoesNotMatchException  = errors.New("Password does not match")
	ERR_InsufficientAuthorityException = errors.New("Insufficient authority.")
	ERR_NoAnonymousAllowedException    = errors.New("No anonymous allowed.")
)

type ErrorList struct {
	ERR_ConflictException              IError
	ERR_ValidationExeption             IError
	ERR_MalformedEntityException       IError
	ERR_UserDoesNotExistException      IError
	ERR_UserEmailAlreadyExistException IError
	ERR_UserAlreadyExistException      IError
	ERR_PasswordDoesNotMatchException  IError
	ERR_InsufficientAuthorityException IError
	ERR_NoAnonymousAllowedException    IError
}

func ThrowError() ErrorList {
	return ErrorList{
		ERR_ConflictException:              ERR_ConflictException,
		ERR_ValidationExeption:             ERR_ValidationExeption,
		ERR_MalformedEntityException:       ERR_MalformedEntityException,
		ERR_UserDoesNotExistException:      ERR_UserDoesNotExistException,
		ERR_UserEmailAlreadyExistException: ERR_UserEmailAlreadyExistException,
		ERR_UserAlreadyExistException:      ERR_UserAlreadyExistException,
		ERR_PasswordDoesNotMatchException:  ERR_PasswordDoesNotMatchException,
		ERR_InsufficientAuthorityException: ERR_InsufficientAuthorityException,
		ERR_NoAnonymousAllowedException:    ERR_NoAnonymousAllowedException,
	}
}
