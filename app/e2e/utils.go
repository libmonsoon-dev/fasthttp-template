package e2e

import "github.com/libmonsoon-dev/fasthttp-template/app"

func must(err error)  {
	if err != nil {
		panic(err)
	}
}

func mustApp(app app.App, err error) app.App {
	must(err)
	return app
}