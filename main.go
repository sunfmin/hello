package main

import (
	"net/http"

	"github.com/sunfmin/hello/config"
	"github.com/sunfmin/hello/handlers"
	aklog "github.com/theplant/appkit/log"
	"github.com/theplant/appkit/server"
)

func main() {
	m := http.NewServeMux()

	mani := config.MustGetManifest()
	mani.Mount(m)

	m.HandleFunc("/", handlers.Home)

	l := aklog.Default()
	h := server.DefaultMiddleware(l)(m)
	l.Info().Log("msg", "Starting at 8080")
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		l.WrapError(err).Log()
	}
}
