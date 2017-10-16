package handlers

import (
	"fmt"
	"net/http"

	"github.com/sunfmin/hello/config"
	"github.com/sunfmin/hello/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	body := templates.Layout(config.MustGetManifest(), templates.Home())
	fmt.Fprint(w, body)
}
