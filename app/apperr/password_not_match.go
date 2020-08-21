package apperr

import (
	"fmt"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"net/http"
)

func NewPasswordNotMatch(err error) PasswordNotMatch {
	return PasswordNotMatch{err}
}

var ErrPasswordNotMatch = PasswordNotMatch{}

const passwordNotMatchMessage = "login/password not match"

type PasswordNotMatch struct {
	err error
}

func (p PasswordNotMatch) HttpStatus() int {
	return http.StatusUnauthorized
}

func (p PasswordNotMatch) PublicError() (code app.ErrorCode, publicMessage string) {
	return app.CodePasswordNotMatch, passwordNotMatchMessage
}

func (p PasswordNotMatch) Error() string {
	if p.err != nil {
		return fmt.Sprintf("%v: %v", passwordNotMatchMessage, p.err.Error())
	}
	return passwordNotMatchMessage
}

func (p PasswordNotMatch) Unwrap() error {
	return p.err
}

var _ app.Error = new(PasswordNotMatch)
