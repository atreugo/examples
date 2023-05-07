package main

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v4"
)

var httpRespBody string = `
<h1>Atreugo Static File Example</h1>
<h2>Links are ordered by func main's code flow.</h2>

<p>Directory or file mapped directly</p>
<a href="/main">Directory mapped at /main</a><br/>
<a href="/readme">README.md on /readme</a></br>
<a href="/gitignore">.gitignore on /gitignore, with middlewares configured</a></br>

<p>Static group with prefix /static</p>
<a href="/static/default/">To /static/default</a><br/>
<a href="/static/middlewares/">To /static/middlewares, with middlewares configured</a><br/>
<a href="/static/custom">You will see forbidden here on /static/custom</a><br/>
<a href="/static/readme">README.md on /static/readme</a></br>
<a href="/static/gitignore">.gitignore on /static/gitignore, with middlewares configured</a></br>
`

func main() {
	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)

	// Register before middlewares
	server.UseBefore(beforeGlobal)

	// Register after middlewares
	server.UseAfter(afterGlobal)

	// Middlewares collection
	middlewares := atreugo.Middlewares{
		Before: []atreugo.Middleware{beforeView},
		After:  []atreugo.Middleware{afterView},
	}

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		defer logger.Info("\t\tview's defer")

		return ctx.HTTPResponse(httpRespBody, 200)
	}).Middlewares(middlewares)

	// Serve files with default configuration
	server.Static("/main", "./")

	// Serve just one file
	server.ServeFile("/readme", "README.md")

	// Serve just one file with middlewares
	server.ServeFile("/gitignore", ".gitignore").Middlewares(middlewares)

	// Creates a new group to serve static files
	static := server.NewGroupPath("/static")

	// Serves files with default configuration
	static.Static("/default", "./")

	// Serves files with default configuration and middlewares
	static.Static("/middlewares", "./").Middlewares(middlewares)

	// Serves files with your own custom configuration
	static.StaticCustom("/custom", &atreugo.StaticFS{
		Root:               "./",
		GenerateIndexPages: false,
		AcceptByteRange:    false,
		Compress:           true,
	}).SkipMiddlewares(beforeGlobal)

	// Serve just one file
	static.ServeFile("/readme", "README.md").UseBefore(beforeView)

	// Serve just one file with middlewares
	static.ServeFile("/gitignore", ".gitignore").Middlewares(middlewares)

	// Run
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
