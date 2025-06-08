package main

import (
	"context"
	"net/http"

	handlerhttp "github.com/ishanwardhono/transfer-system/internal/handler/http"
	accountrepo "github.com/ishanwardhono/transfer-system/internal/repository/account"
	"github.com/ishanwardhono/transfer-system/internal/repository/dbtrx"
	transactionrepo "github.com/ishanwardhono/transfer-system/internal/repository/transaction"
	accountsvc "github.com/ishanwardhono/transfer-system/internal/service/account"
	transactionsvc "github.com/ishanwardhono/transfer-system/internal/service/transaction"
	"github.com/ishanwardhono/transfer-system/pkg/config"
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

	dbtrxRepo := dbtrx.NewRepository(db)

	accountRepo := accountrepo.NewRepository(db)
	accountService := accountsvc.NewService(accountRepo)

	transactionRepo := transactionrepo.NewRepository(db)
	transactionService := transactionsvc.NewService(dbtrxRepo, transactionRepo, accountRepo)

	httpHandler := handlerhttp.NewHandler(accountService, transactionService)
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
