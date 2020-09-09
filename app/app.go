package app

import (
	"github.com/valyala/fasthttp"
	"net"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

type App struct {
	ctx    *Context
	wg     *sync.WaitGroup
	logger Logger
	server *fasthttp.Server
	Root
}

func NewApp(ctx *Context, wg *sync.WaitGroup, logger Logger, server *fasthttp.Server, root Root) App {
	return App{
		ctx,
		wg,
		logger,
		server,
		root,
	}
}

func (app App) Start() {
	go app.signalListener()
	go app.startServer()

	app.wg.Add(1)
	go func() {
		<-app.ctx.Done()
		app.ShutdownServer()
		app.wg.Done()
	}()

	app.wg.Wait()
}

func (app App) StartForTest(listener net.Listener) {

	go func() {
		if err := app.server.Serve(listener); err != nil {
			app.logger.Printf("%+v", err)
		}
	}()

	for app.server.GetOpenConnectionsCount() == 0 {
		runtime.Gosched()
	}
}

func (app App) signalListener() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	app.logger.Printf("Got signal: %v\n", <-signals)
	app.ctx.cancel()
}

func (app App) startServer() {

	go func() {
		for app.server.GetOpenConnectionsCount() == 0 {
			runtime.Gosched()
		}

		app.logger.Printf("Server started on http://%v", app.Config.ServerAddress)
	}()

	if err := app.server.ListenAndServe(app.Config.ServerAddress); err != nil {
		app.logger.Printf("%+v", err)
	}
}

func (app App) ShutdownServer() {
	if err := app.server.Shutdown(); err != nil {
		app.logger.Printf("%+v", err)
	}

}
