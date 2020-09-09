package e2e

import (
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/libmonsoon-dev/fasthttp-template/app/di"
	"github.com/libmonsoon-dev/fasthttp-template/app/infrastructure/config"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"net"
	"os"
)

const AppUrl = "http://localhost"

func Init() (app.Root, Client, func()) {
	must(os.Setenv(config.ServerAddressKey, AppUrl))
	must(os.Setenv(config.JwtSecretKey, "JWT_SECRET"))

	app := mustApp(di.CreateApp())
	listener := fasthttputil.NewInmemoryListener()
	go app.StartForTest(listener)

	c := &fasthttp.Client{
		Dial: func(string) (net.Conn, error) {
			return listener.Dial()
		},
	}

	return app.Root, NewClient(c), func() {
		app.ShutdownServer()
	}
}
