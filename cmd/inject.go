package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AppContainer struct {
	HTTPServer *chi.Mux
}

func NewAppContainer() *AppContainer {
	httpServer := chi.NewRouter()
	return &AppContainer{
		HTTPServer: httpServer,
	}
}

func (c *AppContainer) RunHTTPServer(address string) error {
	return http.ListenAndServe(address, c.HTTPServer)
}
