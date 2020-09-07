package apperr

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"net/http"
)

type InvalidParams struct {
	err error
}

func (i InvalidParams) Error() string {
	if i.err == nil {
		return "invalid params"
	}

	return fmt.Sprintf("invalid params: %v", i.err.Error())
}

func (i InvalidParams) Unwrap() error {
	return i.err
}

func (i InvalidParams) PublicError() (code app.ErrorCode, publicMessage string) {
	code = app.CodeInvalidParams
	publicMessage = "invalid params"

	if ve, ok := i.err.(validator.ValidationErrors); i.err != nil && ok {
		publicMessage = "invalid params: " + ve.Error()
	}
	return
}

func (i InvalidParams) HttpStatus() int {
	return http.StatusBadRequest
}

func NewInvalidParams(err error) InvalidParams {
	return InvalidParams{err}
}

var _ app.Error = new(InvalidParams)
