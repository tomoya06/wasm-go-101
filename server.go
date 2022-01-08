package main

import (
	"log"
	"net/http"
)

const (
	AddSrv       = ":8080"
	TemplatesDir = "."
)

func main() {
	fileSvr := http.FileServer(http.Dir(TemplatesDir))

	if err := http.ListenAndServe(AddSrv, fileSvr); err != nil {
		log.Fatalln(err)
	}
}
