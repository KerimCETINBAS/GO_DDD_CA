package register_command

import (
	_ "gopkg.in/go-playground/validator.v9"
)

type RegisterUserCommand struct {
	UserName string `validate:"required"`
	Password []byte `validate:"required"`
	Email    string `validate:"emai"`
}
