package service

import (
	"context"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/libmonsoon-dev/fasthttp-template/app/apperr"
	"github.com/libmonsoon-dev/fasthttp-template/app/domain"
	"github.com/pkg/errors"
)

type UserService struct {
	logger         app.Logger
	userRepository app.UserRepository
}

func NewUserService(logger app.Logger, userRepository app.UserRepository) *UserService {
	return &UserService{
		logger,
		userRepository,
	}
}

func (us UserService) Create(ctx context.Context, email string, password []byte) (id int, err error) {
	user := domain.User{
		Email:        email,
		PasswordHash: us.generatePasswordHash(password),
	}

	return us.userRepository.Store(ctx, user)
}

func (us UserService) FindById(ctx context.Context, id int) (user domain.User, err error) {
	return us.userRepository.FindById(ctx, id)
}

func (us UserService) FindByEmailPass(ctx context.Context, email string, password []byte) (domain.User, error) {
	user, err := us.FindByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, apperr.ErrItemNotFound) {
			return domain.User{}, apperr.NewPasswordNotMatch(err)
		}
		return domain.User{}, err
	}

	if !us.comparePassword(password, user.PasswordHash) {
		return domain.User{}, errors.WithStack(apperr.ErrPasswordNotMatch)
	}

	return user, nil
}

func (us UserService) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	return us.userRepository.FindByEmail(ctx, email)
}
