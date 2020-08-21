package apperr

import (
	"fmt"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"net/http"
)

const itemNotFoundMessage = "Item not found"

type ItemNotFound struct {
	err error
}

func (err ItemNotFound) Error() string {
	if err.err != nil {
		return fmt.Sprintf("%v: %v", itemNotFoundMessage, err.err.Error())
	}
	return itemNotFoundMessage
}

func (err ItemNotFound) Unwrap() error {
	return err.err
}

func (err ItemNotFound) PublicError() (code app.ErrorCode, publicMessage string) {
	return app.CodeItemNotFound, itemNotFoundMessage
}

func (err ItemNotFound) HttpStatus() int {
	return http.StatusNotFound
}

var ErrItemNotFound = ItemNotFound{}

var _ app.Error = new(ItemNotFound)
