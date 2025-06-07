package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ishanwardhono/transfer-system/config"
	handlerhttp "github.com/ishanwardhono/transfer-system/internal/handler/http"
	"github.com/ishanwardhono/transfer-system/pkg/db"
)

type AppContainer struct {
	HTTPServer *chi.Mux
}

func NewAppContainer(cfg *config.Config) (*AppContainer, error) {
	httpServer := chi.NewRouter()

	_, err := db.NewDatabase(cfg.Database)
	if err != nil {
		return nil, err
	}

	handlerhttp.SetupRouter(httpServer)

	return &AppContainer{
		HTTPServer: httpServer,
	}, nil
}

func (c *AppContainer) RunHTTPServer(address string) error {
	return http.ListenAndServe(address, c.HTTPServer)
}
