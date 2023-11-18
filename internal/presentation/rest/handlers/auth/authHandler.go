package auth_handler

import (
	"errors"

	auth_contract "gihub.com/kerimcetinbas/go_ddd_ca/Contracts/auth"
	register_command "gihub.com/kerimcetinbas/go_ddd_ca/application/authentication/commands/register"
	login_query "gihub.com/kerimcetinbas/go_ddd_ca/application/authentication/queries/login"
	domain_errors "gihub.com/kerimcetinbas/go_ddd_ca/domain/common/Errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/dig"
)

type IAuthHandler interface {
	Login(*fiber.Ctx) error
	Register(*fiber.Ctx) error
}

type authHandler struct{}

type authRouterOpts struct {
	dig.In
	App     *fiber.App
	Handler IAuthHandler
}

var AuthRouter = func(authHandler authRouterOpts) {
	r := authHandler.App.Group("/auth")

	r.Post("/register", authHandler.Handler.Register)
	r.Post("/login", authHandler.Handler.Login)
}

func AuthHandler() IAuthHandler {
	return &authHandler{}
}

// Login implements IAuthHandler.
func (*authHandler) Login(ctx *fiber.Ctx) error {
	body := &auth_contract.AuthLoginRequest{}
	ctx.BodyParser(body)
	query := &login_query.LoginUserQuery{
		Email:    body.Email,
		Password: []byte(body.Password),
	}
	r, err := mediatr.Send[*login_query.LoginUserQuery, *login_query.LoginUserQueryResponse](ctx.Context(), query)
	if err != nil {
		if errors.Is(err, domain_errors.ERR_UserDoesNotExistException) {
			return fiber.ErrUnauthorized

		} else if errors.Is(err, domain_errors.ERR_PasswordDoesNotMatchException) {
			return fiber.ErrUnauthorized
		} else {
			return fiber.ErrInternalServerError
		}
	}
	return ctx.JSON(r)

}

func (*authHandler) Register(ctx *fiber.Ctx) error {
	body := &auth_contract.AuthRegisterRequest{}
	ctx.BodyParser(body)

	command := &register_command.RegisterUserCommand{
		Email:    body.Email,
		Password: []byte(body.Password),
		UserName: body.UserName,
	}

	r, err := mediatr.Send[*register_command.RegisterUserCommand, *register_command.RegisterUserCommandResponse](ctx.Context(), command)

	if err != nil {
		ctx.Status(409)
		return ctx.JSON(fiber.Map{
			"type":   "https://example.com/probs/cant-view-account-details",
			"title":  "User already exist",
			"status": 409,

			"detail":   "User already exist on system",
			"instance": ctx.Request().URI().Path(),
		}, "application/problem+json; charset=utf-8")
	}
	return ctx.JSON(r)
}
