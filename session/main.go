package main

import (
	"flag"
	"log"
	"time"

	"github.com/atreugo/session"
	"github.com/atreugo/session/providers/memcache"
	"github.com/atreugo/session/providers/memory"
	"github.com/atreugo/session/providers/mysql"
	"github.com/atreugo/session/providers/postgre"
	"github.com/atreugo/session/providers/redis"
	"github.com/atreugo/session/providers/sqlite3"
	"github.com/savsgio/atreugo/v11"
)

const defaultProvider = "memory"

var serverSession *session.Session

func init() {
	providerName := flag.String("provider", defaultProvider, "Name of provider")
	flag.Parse()

	var provider session.Provider
	var err error

	switch *providerName {
	case "memory":
		provider, err = memory.New(memory.Config{})
	case "redis":
		// encoder = session.MSGPEncode
		// decoder = session.MSGPDecode
		provider, err = redis.New(redis.Config{
			KeyPrefix:   "session",
			Addr:        "127.0.0.1:6379",
			PoolSize:    8,
			IdleTimeout: 30 * time.Second,
		})
	case "memcache":
		// encoder = session.MSGPEncode
		// decoder = session.MSGPDecode
		provider, err = memcache.New(memcache.Config{
			KeyPrefix: "session",
			ServerList: []string{
				"0.0.0.0:11211",
			},
			MaxIdleConns: 8,
		})

	case "mysql":
		cfg := mysql.NewConfigWith("127.0.0.1", 3306, "root", "session", "test", "session")
		provider, err = mysql.New(cfg)
	case "postgre":
		cfg := postgre.NewConfigWith("127.0.0.1", 5432, "postgres", "session", "test", "session")
		provider, err = postgre.New(cfg)
	case "sqlite3":
		cfg := sqlite3.NewConfigWith("test.db", "session")
		provider, err = sqlite3.New(cfg)
	default:
		panic("Invalid provider")
	}

	if err != nil {
		log.Fatal(err)
	}

	cfg := session.NewDefaultConfig()
	// cfg.EncodeFunc = encoder
	// cfg.DecodeFunc = decoder
	serverSession = session.New(cfg)

	if err = serverSession.SetProvider(provider); err != nil {
		log.Fatal(err)
	}

	log.Print("Starting example with provider: " + *providerName)
}

func main() {
	server := atreugo.New(atreugo.Config{
		Addr: "0.0.0.0:8086",
	})
	server.GET("/", indexHandler)
	server.GET("/set", setHandler)
	server.GET("/get", getHandler)
	server.GET("/delete", deleteHandler)
	server.GET("/getAll", getAllHandler)
	server.GET("/flush", flushHandler)
	server.GET("/destroy", destroyHandler)
	server.GET("/sessionid", sessionIDHandler)
	server.GET("/regenerate", regenerateHandler)
	server.GET("/setexpiration", setExpirationHandler)
	server.GET("/getexpiration", getExpirationHandler)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
