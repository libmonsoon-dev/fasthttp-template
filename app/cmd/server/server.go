package main

import (
	"fmt"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/libmonsoon-dev/fasthttp-template/app/di"
)

func must(app app.App, err error) app.App  {
	if err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}
	return app
}

func main() {
	must(di.CreateApp()).Start()
}
