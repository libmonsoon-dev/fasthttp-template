package service

import (
	"context"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/libmonsoon-dev/fasthttp-template/app/domain"
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

func (us UserService) Create(ctx context.Context, user domain.User) (id int, err error) {
	return us.userRepository.Store(ctx, user)
}

func (us UserService) FindById(ctx context.Context, id int) (user domain.User, err error) {
	return us.userRepository.FindById(ctx, id)
}

func (us UserService) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	return us.userRepository.FindByEmail(ctx, email)
}
