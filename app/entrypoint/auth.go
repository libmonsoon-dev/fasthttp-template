package entrypoint

import (
	"context"
	"encoding/base64"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/libmonsoon-dev/fasthttp-template/app/domain"
	"github.com/libmonsoon-dev/fasthttp-template/app/service"
)

func NewAuthEntrypoint(logger app.Logger, service *service.AuthService) *AuthEntrypoint {
	return &AuthEntrypoint{
		logger,
		service,
	}
}

type AuthEntrypoint struct {
	logger  app.Logger
	service *service.AuthService
}

func (e AuthEntrypoint) SignUp(ctx context.Context, params domain.SignInParams) (domain.AuthToken, error) {
	panic("implement me")
	password, err := base64.StdEncoding.DecodeString(params.Base64Password)
	if err != nil {
		// TODO
	}
	return e.service.SignUp(ctx, params.Email, password)
}

func (e AuthEntrypoint) SignIn(ctx context.Context, params domain.SignInParams) (domain.AuthToken, error) {
	panic("implement me")
	password, err := base64.StdEncoding.DecodeString(params.Base64Password)
	if err != nil {
		// TODO
	}
	return e.service.SignIn(ctx, params.Email, password)
}
