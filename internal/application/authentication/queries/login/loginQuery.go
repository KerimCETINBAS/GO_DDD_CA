package login_query

type LoginUserQuery struct {
	Email    string `validate:"required,email"`
	Password []byte `validate:"required"`
}
