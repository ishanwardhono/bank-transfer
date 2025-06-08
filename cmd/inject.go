package main

import (
	"context"
	"net/http"

	"github.com/ishanwardhono/transfer-system/config"
	handlerhttp "github.com/ishanwardhono/transfer-system/internal/handler/http"
	accountrepo "github.com/ishanwardhono/transfer-system/internal/repository/account"
	accountsvc "github.com/ishanwardhono/transfer-system/internal/service/account"
	"github.com/ishanwardhono/transfer-system/pkg/db"
)

type AppContainer struct {
	HTTPServer *http.Server
}

func NewAppContainer(cfg *config.Config) (*AppContainer, error) {
	db, err := db.NewDatabase(cfg.Database)
	if err != nil {
		return nil, err
	}

	accountRepo := accountrepo.NewRepository(db)
	accountService := accountsvc.NewService(accountRepo)

	httpHandler := handlerhttp.NewHandler(accountService)
	httpRouter := handlerhttp.SetupRouter(httpHandler)
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
