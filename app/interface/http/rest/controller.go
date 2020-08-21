package rest

import (
	"encoding/json"
	"github.com/qiangxue/fasthttp-routing"
)

type Controller struct {
	auth *AuthController
}

func NewController(auth *AuthController) *Controller {
	return &Controller{
		auth,
	}
}

func (c Controller) ApplyTo(restRouteGroup *routing.RouteGroup) {
	restRouteGroup.Use(setJsonSerializerMiddleware)

	c.auth.ApplyTo(restRouteGroup.Group("/auth"))
}

func setJsonSerializerMiddleware(context *routing.Context) error {
	context.Serialize = json.Marshal
	return nil
}