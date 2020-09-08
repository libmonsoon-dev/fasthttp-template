package dto

import "github.com/libmonsoon-dev/fasthttp-template/app/domain"

type AuthToken struct {
	UserId int `json:"userId"`
}

func AuthTokenFrom(model domain.AuthToken) AuthToken {
	return AuthToken(model)
}
