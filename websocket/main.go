package main

import (
	"log"

	"github.com/atreugo/websocket"
	"github.com/savsgio/atreugo/v11"
)

var upgrader = websocket.New(websocket.Config{
	AllowedOrigins: []string{"*"},
})

func main() {
	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)

	server.UseBefore(func(ctx *atreugo.RequestCtx) error {
		ctx.SetUserValue("name", "atreugo")

		return ctx.Next()
	})

	server.GET("/ws", upgrader.Upgrade(func(ws *websocket.Conn) error {
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				return err
			}

			name := ws.UserValue("name").(string)

			log.Printf("User value: %s", name)
			log.Printf("recv: %s", message)

			err = ws.WriteMessage(mt, append(message, " - "+name...))
			if err != nil {
				return err
			}
		}
	}))

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
