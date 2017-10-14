package main

import (
	"fmt"
	"net/http"

	aklog "github.com/theplant/appkit/log"
	"github.com/theplant/appkit/server"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, V3 with log")
	})

	l := aklog.Default()
	h := server.DefaultMiddleware(l)(m)
	l.Info().Log("msg", "Starting at 8080")
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		l.WrapError(err).Log()
	}
}
