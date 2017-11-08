package main

import (
	"golang-starter-kit/route"
	"net/http"
)

func main() {
	r := route.Router()

	server := &http.Server{
		Handler: r,
		Addr:    ":80",
	}
	server.ListenAndServe()
}
