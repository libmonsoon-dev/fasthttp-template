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

func TestApp() (*fasthttp.Client, func())  {
	os.Setenv(config.ServerAddressKey, AppUrl)

	app, err := di.CreateApp()
	if err != nil {
		panic(err)
	}
	listener := fasthttputil.NewInmemoryListener()
	go app.StartForTest(listener)
	client := &fasthttp.Client{
		Dial: func(string) (net.Conn, error) {
			return listener.Dial()
		},
	}

	return client, func() {
		go app.ShutdownServer()
	}
}
