package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/libmonsoon-dev/fasthttp-template/app/apperr"
	"github.com/pkg/errors"
	"sync"
	"time"
)

type AuthService struct {
	logger      app.Logger
	userService app.UserService
	config      *app.Config
	*jwtMapClaimsPool
}

const (
	jwtUserIdKey = "userId"
	jwtExpKey    = "exp"
)

func NewAuthService(logger app.Logger, userService app.UserService, config *app.Config) *AuthService {
	return &AuthService{
		logger,
		userService,
		config,
		newJwtMapClaimsPool(),
	}
}

func (as AuthService) SignUp(ctx context.Context, email string, password []byte) (string, error) {
	id, err := as.userService.Create(ctx, email, password)
	if err != nil {
		return "", err
	}

	return as.EncodeAuthToken(id)
}

func (as AuthService) SignIn(ctx context.Context, email string, password []byte) (string, error) {
	user, err := as.userService.FindByEmailPass(ctx, email, password)

	if err != nil {
		return "", err
	}

	return as.EncodeAuthToken(user.ID)
}

func (as AuthService) EncodeAuthToken(userId int) (token string, err error) {
	claims := as.jwtMapClaimsPool.Acquire()
	defer as.jwtMapClaimsPool.Release(claims)

	(*claims)[jwtUserIdKey] = userId
	(*claims)[jwtExpKey] = time.Now().Add(time.Minute * 15).Unix()

	token, err = jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(as.config.JWTSecret)

	if err != nil {
		return "", errors.WithStack(err)
	}
	return token, nil

}

func (as AuthService) DecodeAuthToken(token string) (userId int, err error) {
	tokenKeyFunc := func(token *jwt.Token) (interface{}, error) {
		return as.config.JWTSecret, nil
	}

	claims := as.jwtMapClaimsPool.Acquire()
	defer as.jwtMapClaimsPool.Release(claims)

	if _, err := jwt.ParseWithClaims(token, claims, tokenKeyFunc); err != nil {
		return 0, errors.WithStack(err)
	}

	if userIdFloat, ok := (*claims)[jwtUserIdKey].(float64); !ok {
		return 0, errors.WithStack(apperr.NewInternalError(fmt.Errorf("error parsing userId")))
	} else {
		userId = int(userIdFloat)
	}

	return
}

func newJwtMapClaimsPool() *jwtMapClaimsPool {
	return &jwtMapClaimsPool{
		sync.Pool{
			New: func() interface{} {
				return &jwt.MapClaims{}
			},
		},
	}
}

type jwtMapClaimsPool struct {
	pool sync.Pool
}

func (p *jwtMapClaimsPool) Acquire() *jwt.MapClaims {
	return p.pool.Get().(*jwt.MapClaims)
}

func (p *jwtMapClaimsPool) Release(claims *jwt.MapClaims) {
	for key := range *claims {
		delete(*claims, key)
	}
	p.pool.Put(claims)
}
