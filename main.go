package main

import (
	"net/http"
	"zeroweb/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    "localhost:8080",
	}
	server.ListenAndServe()
}
