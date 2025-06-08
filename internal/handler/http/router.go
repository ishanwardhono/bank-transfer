package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ishanwardhono/transfer-system/internal/service/account"
	"github.com/ishanwardhono/transfer-system/internal/service/transaction"
	"github.com/ishanwardhono/transfer-system/pkg/httphelper"
)

func SetupRouter(handler *Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(httphelper.LoggerMiddleware)

	router.Post("/accounts", handler.RegisterAccount)
	router.Get("/accounts/{accountId}", handler.GetAccountById)
	router.Post("/transactions", handler.Transfer)

	return router
}

type Handler struct {
	accountService  account.Service
	transferService transaction.Service
}

func NewHandler(
	accountService account.Service,
	transferService transaction.Service,
) *Handler {
	return &Handler{
		accountService:  accountService,
		transferService: transferService,
	}
}
