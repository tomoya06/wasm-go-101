package main

import (
	"log"
	"net/http"
	"strings"
)

const (
	AddSrv       = ":8080"
	TemplatesDir = "."
)

const dir = "."

func main() {
	fs := http.FileServer(http.Dir(dir))
	log.Print("Serving " + dir + " on http://localhost" + AddSrv)
	http.ListenAndServe(AddSrv, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")

		// Warning!! without this the wasm file will not be executed
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	}))
}
