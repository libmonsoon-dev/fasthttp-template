package e2e

import (
	"github.com/libmonsoon-dev/fasthttp-template/app/di"
	"github.com/libmonsoon-dev/fasthttp-template/app/infrastructure/config"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"net"
	"os"
)

const AppUrl = "http://localhost"

func Init() (Client, func()) {
	must(os.Setenv(config.ServerAddressKey, AppUrl))

	app := mustApp(di.CreateApp())
	listener := fasthttputil.NewInmemoryListener()
	go app.StartForTest(listener)

	c := &fasthttp.Client{
		Dial: func(string) (net.Conn, error) {
			return listener.Dial()
		},
	}

	return NewClient(c), func() { app.ShutdownServer() }
}
