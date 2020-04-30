package main

import (
	"fmt"
	"time"

	"github.com/savsgio/atreugo/v11"
)

func indexHandler(ctx *atreugo.RequestCtx) error {
	html := "<h2>Welcome to use session, you should request to the: </h2>"

	html += `> <a href="/">/</a><br>`
	html += `> <a href="/set">set</a><br>`
	html += `> <a href="/get">get</a><br>`
	html += `> <a href="/delete">delete</a><br>`
	html += `> <a href="/getAll">getAll</a><br>`
	html += `> <a href="/flush">flush</a><br>`
	html += `> <a href="/destroy">destroy</a><br>`
	html += `> <a href="/sessionid">sessionid</a><br>`
	html += `> <a href="/regenerate">regenerate</a><br>`

	return ctx.HTTPResponse(html)
}

func setHandler(ctx *atreugo.RequestCtx) (err error) {
	store, err := serverSession.Get(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = serverSession.Save(ctx, store)
	}()

	store.Set("foo", "bar")

	return ctx.TextResponse(fmt.Sprintf("Session SET: foo='%s' --> OK", store.Get("foo").(string)))
}

func getHandler(ctx *atreugo.RequestCtx) (err error) {
	store, err := serverSession.Get(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = serverSession.Save(ctx, store)
	}()

	val := store.Get("foo")
	if val == nil {
		return ctx.TextResponse("Session GET: foo is nil")
	}

	return ctx.TextResponse(fmt.Sprintf("Session GET: foo='%s'", val.(string)))
}

func deleteHandler(ctx *atreugo.RequestCtx) (err error) {
	store, err := serverSession.Get(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = serverSession.Save(ctx, store)
	}()

	store.Delete("foo")

	val := store.Get("name")
	if val == nil {
		return ctx.TextResponse("Session DELETE: foo --> OK")
	}

	return ctx.TextResponse("Session DELETE: foo --> ERROR")
}

func getAllHandler(ctx *atreugo.RequestCtx) (err error) {
	store, err := serverSession.Get(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = serverSession.Save(ctx, store)
	}()

	store.Set("foo1", "bar1")
	store.Set("foo2", 2)
	store.Set("foo3", "bar3")
	store.Set("foo4", []byte("bar4"))

	data := store.GetAll()

	fmt.Println(data)

	return ctx.TextResponse("Session GetAll: See the OS console!")
}

func flushHandler(ctx *atreugo.RequestCtx) (err error) {
	store, err := serverSession.Get(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = serverSession.Save(ctx, store)
	}()

	store.Flush()

	data := store.GetAll()

	fmt.Println(data)

	return ctx.TextResponse("Session FLUSH: See the OS console!")
}

func destroyHandler(ctx *atreugo.RequestCtx) error {
	err := serverSession.Destroy(ctx)
	if err != nil {
		return err
	}

	return ctx.TextResponse("Session DESTROY --> OK")
}

func sessionIDHandler(ctx *atreugo.RequestCtx) (err error) {
	store, err := serverSession.Get(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = serverSession.Save(ctx, store)
	}()

	sessionID := store.GetSessionID()
	ctx.SetBodyString("Session: Current session id: ")
	ctx.Write(sessionID)

	return nil
}

func regenerateHandler(ctx *atreugo.RequestCtx) error {
	if err := serverSession.Regenerate(ctx); err != nil {
		return err
	}

	store, err := serverSession.Get(ctx)
	if err != nil {
		return err
	}

	ctx.SetBodyString("Session REGENERATE: New session id: ")
	ctx.Write(store.GetSessionID())

	return nil
}

func getExpirationHandler(ctx *atreugo.RequestCtx) error {
	store, err := serverSession.Get(ctx)
	if err != nil {
		return err
	}

	expiration := store.GetExpiration()

	ctx.SetBodyString("Session Expiration: ")
	ctx.WriteString(expiration.String())

	return nil
}

func setExpirationHandler(ctx *atreugo.RequestCtx) (err error) {
	store, err := serverSession.Get(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = serverSession.Save(ctx, store)
	}()

	err = store.SetExpiration(30 * time.Second)
	if err != nil {
		return err
	}

	return ctx.TextResponse("Session Expiration set to 30 seconds")
}
