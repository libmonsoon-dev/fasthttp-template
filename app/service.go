package app

import (
	"context"
	"github.com/libmonsoon-dev/fasthttp-template/app/domain"
)

type UserService interface {
	Create(ctx context.Context, email string, password []byte) (id int, err error)
	FindById(ctx context.Context, id int) (user domain.User, err error)
	FindByEmailPass(ctx context.Context, email string, password []byte) (user domain.User, err error)
	FindByEmail(ctx context.Context, email string) (user domain.User, err error)
}

type AuthService interface {
	SignUp(ctx context.Context, email string, password []byte) (string, error)
	SignIn(ctx context.Context, email string, password []byte) (string, error)
	EncodeAuthToken(userId int) (token string, err error)
	DecodeAuthToken(token string) (userId int, err error)
}
