package domain

type SignInParams struct {
	Email          string `validate:"required,email"`
	Base64Password string `validate:"required,base64,min=10,max=64"`
}
