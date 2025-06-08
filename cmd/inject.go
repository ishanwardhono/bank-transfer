package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ishanwardhono/transfer-system/config"
	handlerhttp "github.com/ishanwardhono/transfer-system/internal/handler/http"
	"github.com/ishanwardhono/transfer-system/pkg/db"
)

type AppContainer struct {
	HTTPServer *http.Server
}

func NewAppContainer(cfg *config.Config) (*AppContainer, error) {
	httpRouter := chi.NewRouter()

	_, err := db.NewDatabase(cfg.Database)
	if err != nil {
		return nil, err
	}

	handlerhttp.SetupRouter(httpRouter)
	httpServer := &http.Server{
		Addr:    cfg.GetServerAddress(),
		Handler: httpRouter,
	}

	return &AppContainer{
		HTTPServer: httpServer,
	}, nil
}

func (c *AppContainer) Run() error {
	return c.HTTPServer.ListenAndServe()
}

func (c *AppContainer) Shutdown(ctx context.Context) error {
	return c.HTTPServer.Shutdown(ctx)
}
