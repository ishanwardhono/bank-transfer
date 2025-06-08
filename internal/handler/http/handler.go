package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ishanwardhono/transfer-system/internal/entity/dto"
	appcontext "github.com/ishanwardhono/transfer-system/pkg/context"
	"github.com/ishanwardhono/transfer-system/pkg/errors"
	"github.com/ishanwardhono/transfer-system/pkg/httphelper"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
)

func (h *Handler) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterAccountRequest
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Errorf(ctx, "Failed to parse request body, err: %v", err)
		httphelper.HandleError(w, errors.Wrap(http.StatusBadRequest, err))
		return
	}
	defer r.Body.Close()

	ctx = appcontext.SetAccountID(ctx, request.AccountID)
	if err := request.Validate(); err != nil {
		logger.Errorf(ctx, "Failed to validate request body, err: %v", err)
		httphelper.HandleError(w, errors.Wrap(http.StatusBadRequest, err))
		return
	}

	err := h.accountService.Register(ctx, request)
	if err != nil {
		logger.Errorf(ctx, "Failed to register account, err: %v", err)
		httphelper.HandleError(w, err)
		return
	}

	httphelper.HandleCreatedResponse(w, nil)
}

func (h *Handler) GetAccountById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	accountIdStr := chi.URLParam(r, "accountId")
	if accountIdStr == "" {
		logger.Errorf(ctx, "Missing account ID in request")
		httphelper.HandleError(w, errors.New(http.StatusBadRequest, "Missing account ID"))
		return
	}
	accountId, err := strconv.ParseInt(accountIdStr, 10, 64)
	if err != nil {
		logger.Errorf(ctx, "Invalid account ID format: %v", err)
		httphelper.HandleError(w, errors.Wrap(http.StatusBadRequest, err))
		return
	}

	ctx = appcontext.SetAccountID(ctx, accountId)
	resp, err := h.accountService.GetById(ctx, accountId)
	if err != nil {
		logger.Errorf(ctx, "Failed to register account, err: %v", err)
		httphelper.HandleError(w, err)
		return
	}

	httphelper.HandleResponse(w, resp)
}

func (h *Handler) Transfer(w http.ResponseWriter, r *http.Request) {
	var request dto.TransferRequest
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Errorf(ctx, "Failed to parse request body, err: %v", err)
		httphelper.HandleError(w, errors.Wrap(http.StatusBadRequest, err))
		return
	}
	defer r.Body.Close()

	ctx = appcontext.SetAccountID(ctx, request.SourceAccountID)
	if err := request.Validate(); err != nil {
		logger.Errorf(ctx, "Failed to validate request body, err: %v", err)
		httphelper.HandleError(w, errors.Wrap(http.StatusBadRequest, err))
		return
	}

	err := h.transferService.Transfer(ctx, request)
	if err != nil {
		logger.Errorf(ctx, "Failed to process transfer, err: %v", err)
		httphelper.HandleError(w, err)
		return
	}

	httphelper.HandleResponse(w, nil)
}
