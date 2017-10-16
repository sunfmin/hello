package config

import (
	"os"

	"github.com/theplant/createreactappmanifest"
)

var mcache *manifest.Manifest

func MustGetManifest() *manifest.Manifest {
	if mcache != nil {
		return mcache
	}

	isDev := os.Getenv("DEV") == "1"
	mcfg := &manifest.Config{PublicURL: "/hello"}
	if isDev {
		// mcfg.IsDev = true
		mcfg.ManifestDir = "./front/build"
		// mcfg.DevBundleURL = "http://localhost:3000/static/js/bundle.js"
	} else {
		mcfg.IsDev = false
		mcfg.ManifestDir = "/www"
	}
	m, err := manifest.New(mcfg)
	if err != nil {
		panic(err)
	}
	mcache = m
	return m
}
