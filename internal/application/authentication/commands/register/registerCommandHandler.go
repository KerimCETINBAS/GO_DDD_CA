package register_command

import (
	"context"

	"github.com/kerimcetinbas/go_ddd_ca/application/common/persistence"
	"github.com/kerimcetinbas/go_ddd_ca/application/common/services"
	user_domain "github.com/kerimcetinbas/go_ddd_ca/domain/user"
	user_valueobject "github.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
	"github.com/mehdihadeli/go-mediatr"
)

type RegisterUsercommandHandlerType = mediatr.RequestHandler[*RegisterUserCommand, *RegisterUserCommandResponse]

type RegisterUserCommandHandler struct {
	DateTimeProvider services.IDateTimeProvider
	userRespository  persistence.IUserRepository
}

func Provider(
	dateTimeProvider services.IDateTimeProvider,
	userRespository persistence.IUserRepository,
) RegisterUsercommandHandlerType {

	return &RegisterUserCommandHandler{
		DateTimeProvider: dateTimeProvider,
		userRespository:  userRespository,
	}
}

func Invoke(cmd RegisterUsercommandHandlerType) {
	mediatr.RegisterRequestHandler(cmd)
}

func (c *RegisterUserCommandHandler) Handle(ctx context.Context, command *RegisterUserCommand) (*RegisterUserCommandResponse, error) {

	u := user_domain.NewUser(
		command.Email,
		command.UserName,
		command.Password,
		make([]user_valueobject.FriendIdValueObject, 0),
	)

	if err := c.userRespository.Create(u); err != nil {
		return &RegisterUserCommandResponse{}, err
	}

	r := &RegisterUserCommandResponse{
		Id:        u.Id().Value(),
		UserName:  u.UserName(),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
	}

	return r, nil
}
