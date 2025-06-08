package http

import (
	"encoding/json"
	"net/http"

	"github.com/ishanwardhono/transfer-system/internal/entity/dto"
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
