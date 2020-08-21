package http

import (
	"github.com/libmonsoon-dev/fasthttp-template/app/interface/http/rest"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func NewController(rest *rest.Controller) fasthttp.RequestHandler {
	root := routing.New()

	rest.ApplyTo(root.Group("/rest"))

	return root.HandleRequest
}
