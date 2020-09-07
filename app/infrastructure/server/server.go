package server

import (
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/valyala/fasthttp"
)

func New(logger app.Logger, handler fasthttp.RequestHandler) *fasthttp.Server {
	return &fasthttp.Server{
		Handler: handler,
		Logger: logger,

		DisableKeepalive: true,
	}
}