package apperr

import (
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"net/http"
)

func NewInternalError(err error) InternalError {
	return InternalError{err}
}

type InternalError struct {
	err error
}

func (err InternalError) Error() string {
	if err.err != nil {
		return err.err.Error()
	}
	return "empty InternalError"
}

func (err InternalError) Unwrap() error {
	return err.err
}

func (err InternalError) PublicError() (code app.ErrorCode, publicMessage string) {
	return app.CodeInternalError, "Internal error"
}

func (err InternalError) HttpStatus() int {
	return http.StatusInternalServerError
}

var _ app.Error = new(InternalError)