package main

import (
	"github.com/savsgio/atreugo/v11"
)

func main() {
	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Hello World")
	})

	server.GET("/echo/{path:*}", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Echo message: " + ctx.UserValue("path").(string))
	})

	v1 := server.NewGroupPath("/v1")
	v1.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Hello V1 Group")
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
