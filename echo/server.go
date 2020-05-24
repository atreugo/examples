package main

import (
	"fmt"

	"github.com/savsgio/atreugo/v11"
)

// AtHandler Sugar candy type for Atreugo handlers.
type AtHandler func(ctx *atreugo.RequestCtx) error

// Route groups route information.
type Route struct {
	Group    string
	Endpoint string
	Method   string
	Handler  AtHandler
}

// AtServer Type provides a Atreugo HTTP server.
type AtServer struct {
	engine *atreugo.Atreugo
}

// NewAtreugoServer "Constructor".
func NewAtreugoServer(socket string) AtServer {
	config := atreugo.Config{
		Addr: socket,
	}
	return AtServer{
		engine: atreugo.New(config),
	}
}

// Run Method to start server.
func (a AtServer) Run() error {
	return a.engine.ListenAndServe()
}

// RegisterRoute Method registers routes.
func (a AtServer) RegisterRoute(r Route) {
	a.engine.GET(r.Endpoint, echo)
}

// echo Handler for echo server.
func echo(ctx *atreugo.RequestCtx) error {
	msg := ctx.UserValue("echo")
	return ctx.TextResponse(fmt.Sprintf("Message: %s", msg))
}
