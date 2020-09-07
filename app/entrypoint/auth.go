package entrypoint

import (
	"context"
	"encoding/base64"
	"github.com/go-playground/validator/v10"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/libmonsoon-dev/fasthttp-template/app/apperr"
	"github.com/libmonsoon-dev/fasthttp-template/app/domain"
	"github.com/libmonsoon-dev/fasthttp-template/app/service"
	"github.com/pkg/errors"
)

func NewAuthEntrypoint(validator *validator.Validate, logger app.Logger, service *service.AuthService) *AuthEntrypoint {
	return &AuthEntrypoint{
		validator,
		logger,
		service,
	}
}

type AuthEntrypoint struct {
	validator *validator.Validate
	logger  app.Logger
	service *service.AuthService
}

func (e AuthEntrypoint) SignUp(ctx context.Context, params domain.SignInParams) (domain.AuthToken, error) {
	if err := e.validator.Struct(params); err != nil {
		return domain.AuthToken{}, errors.WithStack(apperr.NewInvalidParams(err))
	}

	password, err := base64.StdEncoding.DecodeString(params.Base64Password)
	if err != nil {
		return domain.AuthToken{}, errors.WithStack(apperr.NewInternalError(err))
	}
	return e.service.SignUp(ctx, params.Email, password)
}

func (e AuthEntrypoint) SignIn(ctx context.Context, params domain.SignInParams) (domain.AuthToken, error) {
	if err := e.validator.Struct(params); err != nil {
		return domain.AuthToken{}, errors.WithStack(apperr.NewInvalidParams(err))
	}

	password, err := base64.StdEncoding.DecodeString(params.Base64Password)
	if err != nil {
		return domain.AuthToken{}, errors.WithStack(apperr.NewInternalError(err))
	}
	return e.service.SignIn(ctx, params.Email, password)
}
