package account

import (
	"context"

	"github.com/ishanwardhono/transfer-system/internal/entity/dto"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
)

func (s *service) Register(ctx context.Context, req dto.RegisterAccountRequest) error {
	accountModel := req.ToModel()
	if err := s.AccountRepo.InsertAccount(ctx, accountModel); err != nil {
		logger.Errorf(ctx, "failed to insert account, err: %v", err)
		return err
	}

	return nil
}

func (s *service) GetById(ctx context.Context, accountId int64) (dto.GetAccountByIdResponse, error) {
	account, err := s.AccountRepo.GetAccount(ctx, accountId)
	if err != nil {
		logger.Errorf(ctx, "failed to insert account, err: %v", err)
		return dto.GetAccountByIdResponse{}, err
	}

	resp := dto.FromModelAccount(account)
	return resp, nil
}
