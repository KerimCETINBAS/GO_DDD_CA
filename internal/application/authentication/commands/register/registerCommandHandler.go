package register_command

import (
	"context"

	"gihub.com/kerimcetinbas/go_ddd_ca/application/common/persistence"
	"gihub.com/kerimcetinbas/go_ddd_ca/application/common/services"
	user_domain "gihub.com/kerimcetinbas/go_ddd_ca/domain/user"
	user_valueobject "gihub.com/kerimcetinbas/go_ddd_ca/domain/user/valueObject"
	"github.com/google/uuid"
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
		Id:        uuid.New(),
		UserName:  command.UserName,
		CreatedAt: c.DateTimeProvider.DateTime(),
	}

	return r, nil
}
