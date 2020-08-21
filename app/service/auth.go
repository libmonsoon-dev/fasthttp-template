package service

import (
	"context"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/libmonsoon-dev/fasthttp-template/app/apperr"
	"github.com/libmonsoon-dev/fasthttp-template/app/domain"
	"github.com/pkg/errors"
)

type AuthService struct {
	logger      app.Logger
	userService *UserService
}

func NewAuthService(logger app.Logger, userService *UserService) *AuthService {
	return &AuthService{
		logger,
		userService,
	}
}

func (as AuthService) SignUp(ctx context.Context, email string, password []byte) (domain.AuthToken, error) {
	user := domain.User{
		Email: email,
		PasswordHash: as.generatePasswordHash(password),
	}

	_, err := as.userService.Create(ctx, user)
	if err != nil {
		return domain.AuthToken{}, err
	}

	return as.GetAuthToken(user), nil
}

func (as AuthService) SignIn(ctx context.Context, email string, password []byte) (domain.AuthToken, error) {
	user, err := as.userService.FindByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, apperr.ErrItemNotFound) {
			return domain.AuthToken{}, apperr.NewPasswordNotMatch(err)
		}
		return domain.AuthToken{}, err
	}

	if !as.comparePassword(password, user.PasswordHash) {
		return domain.AuthToken{}, errors.WithStack(apperr.ErrPasswordNotMatch)
	}

	return as.GetAuthToken(user), nil
}

func (as AuthService) GetAuthToken(user domain.User) domain.AuthToken {
	return domain.AuthToken{
		UserId: user.ID,
	}
}
