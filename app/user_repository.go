package app

import (
	"context"
	"github.com/libmonsoon-dev/fasthttp-template/app/domain"
)

type UserRepository interface {
	Store(ctx context.Context, user domain.User) (id int, err error)
	FindById(ctx context.Context, id int) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
}
