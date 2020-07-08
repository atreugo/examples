package main

import (
	"github.com/savsgio/atreugo/v11"
)

func setUpLocalhost(server *atreugo.Atreugo) {
	vHost := server.NewVirtualHost("localhost:8000")

	vHost.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Host: localhost")
	})
}

func setUpExampleDotCom(server *atreugo.Atreugo) {
	vHost := server.NewVirtualHost("example.com:8000")

	vHost.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Host: example.com")
	})
}

func main() {
	server := atreugo.New(atreugo.Config{
		Addr: "0.0.0.0:8000",
	})

	setUpLocalhost(server)
	setUpExampleDotCom(server)

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Host not found", 404)
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
