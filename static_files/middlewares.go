package main

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v4"
)

func beforeGlobal(ctx *atreugo.RequestCtx) error {
	logger.Info("Middleware executed BEFORE GLOBAL.")

	return ctx.Next()
}

func afterGlobal(ctx *atreugo.RequestCtx) error {
	logger.Info("Middleware executed AFTER GLOBAL.")

	return ctx.Next()
}

func beforeView(ctx *atreugo.RequestCtx) error {
	logger.Info("\tMiddleware executed BEFORE SPECIFIC VIEW @ " + string(ctx.Path()))

	return ctx.Next()
}

func afterView(ctx *atreugo.RequestCtx) error {
	logger.Info("\tMiddleware executed AFTER SPECIFIC VIEW @ " + string(ctx.Path()))

	return ctx.Next()
}
