package login_query

import (
	"context"

	"github.com/kerimcetinbas/go_ddd_ca/application/common/interfaces/auth"
	"github.com/kerimcetinbas/go_ddd_ca/application/common/persistence"
	"github.com/kerimcetinbas/go_ddd_ca/application/common/services"
	domain_errors "github.com/kerimcetinbas/go_ddd_ca/domain/common/Errors"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/dig"
)

type LoginUserQueryHandlerType = mediatr.RequestHandler[*LoginUserQuery, *LoginUserQueryResponse]

type LoginUserQueryHandler struct {
	dateTimeProvider services.IDateTimeProvider
	tokenProvider    auth.ITokenProvider
	userRepository   persistence.IUserRepository
}

type LoginUserQueryParams struct {
	dig.In
	DateTimeProvider services.IDateTimeProvider
	TokenGenerator   auth.ITokenProvider
	UserRepository   persistence.IUserRepository
}

func Provider(
	services LoginUserQueryParams,
) LoginUserQueryHandlerType {
	return &LoginUserQueryHandler{
		dateTimeProvider: services.DateTimeProvider,
		tokenProvider:    services.TokenGenerator,
		userRepository:   services.UserRepository,
	}
}

func Invoke(cmd LoginUserQueryHandlerType) {
	mediatr.RegisterRequestHandler(cmd)
}

func (h *LoginUserQueryHandler) Handle(
	ctx context.Context,
	query *LoginUserQuery) (*LoginUserQueryResponse, error) {

	user, err := h.userRepository.GetByEmail(query.Email)

	if err != nil {

		return &LoginUserQueryResponse{}, err
	}

	if string(user.Password()) != string(query.Password) {
		return &LoginUserQueryResponse{}, domain_errors.ThrowError().ERR_PasswordDoesNotMatchException
	}

	token, err := h.tokenProvider.GenerateAccessToken(user)

	if err != nil {
		return &LoginUserQueryResponse{}, domain_errors.ThrowError().ERR_MalformedEntityException
	}

	return &LoginUserQueryResponse{
		Id:           user.Id().Value(),
		UserName:     user.UserName(),
		Email:        user.Email(),
		AccessToken:  token,
		RefreshToken: "hardcoded, /** todo replace this with actual token",
	}, nil

}
