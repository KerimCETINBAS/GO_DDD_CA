package user_handler

import (
	"github.com/gofiber/fiber/v2"
	users_getall_command "github.com/kerimcetinbas/go_ddd_ca/application/users/queries/getAll"
	"github.com/kerimcetinbas/go_ddd_ca/presentation/rest/middlewares"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/dig"
)

type IUserHandler interface {
	GetAll(*fiber.Ctx) error
}

type userHandler struct{}

type userHandlerOpts struct {
	dig.In
	App            *fiber.App
	Handler        IUserHandler
	AuthMiddleware middlewares.IAuthMiddleware
}

func UserRouter(opts userHandlerOpts) {
	r := opts.App.Group("/users")
	r.Get("/", opts.AuthMiddleware.Validate(), opts.Handler.GetAll)
}

func UserHandler() IUserHandler {
	return &userHandler{}
}

func (*userHandler) GetAll(ctx *fiber.Ctx) error {
	command := &users_getall_command.GetAllUsersQuery{}

	r, _ := mediatr.Send[*users_getall_command.GetAllUsersQuery, *users_getall_command.GetAllUsersQueryResponse](ctx.Context(), command)

	return ctx.JSON(r)
}
