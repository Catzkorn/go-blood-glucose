package main

import (
	"net/http"

	"github.com/Catzkorn/go-blood-glucose/server"
)

func main() {
	http.ListenAndServe(":3000", &server.Server{})
}
