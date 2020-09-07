// +build wireinject

package di

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/libmonsoon-dev/fasthttp-template/app/apperr"
	"github.com/libmonsoon-dev/fasthttp-template/app/domain"
	"github.com/libmonsoon-dev/fasthttp-template/app/entrypoint"
	"github.com/libmonsoon-dev/fasthttp-template/app/infrastructure/config"
	"github.com/libmonsoon-dev/fasthttp-template/app/infrastructure/logger"
	"github.com/libmonsoon-dev/fasthttp-template/app/infrastructure/server"
	"github.com/libmonsoon-dev/fasthttp-template/app/interface/http"
	"github.com/libmonsoon-dev/fasthttp-template/app/interface/http/rest"
	"github.com/libmonsoon-dev/fasthttp-template/app/service"
	"sync"
)

func CreateApp() (app.App, error) {
	panic(wire.Build(
		app.NewApp,
		app.NewContext,
		wire.Value(&sync.WaitGroup{}),
		logger.NewStderrLogger,
		wire.Bind(new(app.Logger), new(*logger.Logger)),
		config.EnvironmentProvider,
		server.New,
		http.NewController,
		rest.NewController,
		rest.NewAuthController,
		validator.New,
		entrypoint.NewAuthEntrypoint,
		service.NewAuthService,
		service.NewUserService,
		newRepo,
		wire.Bind(new(app.UserRepository), new(*repo)),
	))
}

func newRepo() *repo {
	r := make(repo, 0)
	return &r
}

type repo []domain.User

func (r *repo) Store(ctx context.Context, user domain.User) (id int, err error) {
	id = len(*r)
	user.ID = id
	*r = append(*r, user)
	return
}

func (r repo) FindById(ctx context.Context, id int) (domain.User, error) {
	for _, user := range r {
		if user.ID == id {
			return user, nil
		}
	}

	if id < len(r) {
		return r[id], nil
	}
	return domain.User{}, apperr.ErrItemNotFound
}

func (r repo) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	for _, user := range r {
		if user.Email == email {
			return user, nil
		}
	}
	return domain.User{}, apperr.ErrItemNotFound
}