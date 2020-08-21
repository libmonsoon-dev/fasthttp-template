package dto

import "github.com/libmonsoon-dev/fasthttp-template/app/domain"

type SignInParams struct {
	Email          string `json:"email"`
	Base64Password string `json:"password"`
}

func (params SignInParams) Model() domain.SignInParams {
	return domain.SignInParams(params)
}