package login_query

import (
	"context"

	"gihub.com/kerimcetinbas/go_ddd_ca/application/common/interfaces/auth"
	"gihub.com/kerimcetinbas/go_ddd_ca/application/common/persistence"
	"gihub.com/kerimcetinbas/go_ddd_ca/application/common/services"
	domain_errors "gihub.com/kerimcetinbas/go_ddd_ca/domain/common/Errors"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/dig"
)

type LoginUserQueryHandlerType = mediatr.RequestHandler[*LoginUserQuery, *LoginUserQueryResponse]

type LoginUserQueryHandler struct {
	dateTimeProvider services.IDateTimeProvider
	tokenGenerator   auth.IPasetoTokenGenerator
	userRepository   persistence.IUserRepository
}

type LoginUserQueryParams struct {
	dig.In
	DateTimeProvider services.IDateTimeProvider
	TokenGenerator   auth.IPasetoTokenGenerator
	UserRepository   persistence.IUserRepository
}

func Provider(
	services LoginUserQueryParams,
) LoginUserQueryHandlerType {
	return &LoginUserQueryHandler{
		dateTimeProvider: services.DateTimeProvider,
		tokenGenerator:   services.TokenGenerator,
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

	token, err := h.tokenGenerator.GenerateAccessToken(user)

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
