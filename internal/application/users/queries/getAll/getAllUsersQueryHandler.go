package users_getall_query

import (
	"context"

	"gihub.com/kerimcetinbas/go_ddd_ca/application/common/persistence"
	"github.com/mehdihadeli/go-mediatr"
)

type GetAllUserQueryHandler struct {
	userRepository persistence.IUserRepository
}

type GetAllUsersQueryHandlerType = mediatr.RequestHandler[*GetAllUsersQuery, *GetAllUsersQueryResponse]

func Provider(
	userRepository persistence.IUserRepository,
) GetAllUsersQueryHandlerType {
	return &GetAllUserQueryHandler{
		userRepository: userRepository,
	}
}

func Invoke(query GetAllUsersQueryHandlerType) {
	mediatr.RegisterRequestHandler(query)
}

func (q *GetAllUserQueryHandler) Handle(ctx context.Context, command *GetAllUsersQuery) (*GetAllUsersQueryResponse, error) {

	users := make(GetAllUsersQueryResponse, 0)

	us, _ := q.userRepository.GetAll()
	for _, u := range us {
		users = append(users, GetUserQueryResponse{
			Id:       u.Id().Value().String(),
			Email:    u.Email(),
			UserName: u.UserName(),
		})
	}

	return &users, nil
}
