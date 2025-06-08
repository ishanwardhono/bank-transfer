package account

import (
	"context"

	"github.com/ishanwardhono/transfer-system/internal/entity/dto"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
)

func (s *service) Register(ctx context.Context, req dto.RegisterAccountRequest) (dto.RegisterAccountResponse, error) {
	accountModel := req.ToModel()
	if err := s.AccountRepo.InsertAccount(ctx, accountModel); err != nil {
		logger.Errorf(ctx, "failed to insert account, err: %v", err)
		return dto.RegisterAccountResponse{}, err
	}

	return dto.RegisterAccountResponse{
		ID: req.AccountID,
	}, nil
}
