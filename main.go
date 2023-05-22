package main

import (
	"net/http"

	"github.com/j4kubM/basketball-api/controller"
)

func main() {
	handler := controller.New()

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	server.ListenAndServe()
}
