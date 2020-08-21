package rest

import (
	"encoding/json"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/libmonsoon-dev/fasthttp-template/app/entrypoint"
	"github.com/libmonsoon-dev/fasthttp-template/app/interface/http/rest/dto"
	"github.com/pkg/errors"
	"github.com/qiangxue/fasthttp-routing"
)

func NewAuthController(logger app.Logger, entrypoint *entrypoint.AuthEntrypoint) *AuthController {
	return &AuthController{
		logger,
		entrypoint,
	}
}

type AuthController struct {
	logger     app.Logger
	entrypoint *entrypoint.AuthEntrypoint
}

func (ac AuthController) ApplyTo(authRouteGroup *routing.RouteGroup) {
	authRouteGroup.Post("/sign-up", ac.SignUp)
	authRouteGroup.Post("/sign-in", ac.SignIn)
}

func (ac AuthController) SignUp(ctx *routing.Context) error {
	var params dto.SignInParams

	reqBody := ctx.PostBody()
	err := json.Unmarshal(reqBody, &params)
	if err != nil {
		err = errors.WithStack(err)
		ac.logger.Printf("%+v\n", err)
		return err
	}

	result, err := ac.entrypoint.SignUp(ctx, params.Model())
	if err != nil {
		ac.logger.Printf("%+v\n", err)
		return err
	}

	return ctx.WriteData(result)

}

func (ac AuthController) SignIn(ctx *routing.Context) error {
	var params dto.SignInParams

	err := json.Unmarshal(ctx.PostBody(), &params)
	if err != nil {
		ac.logger.Printf("%+v\n", errors.WithStack(err))
		return err
	}

	result, err := ac.entrypoint.SignIn(ctx, params.Model())
	if err != nil {
		ac.logger.Printf("%+v\n", err)
		return err
	}

	return ctx.WriteData(result)

}
