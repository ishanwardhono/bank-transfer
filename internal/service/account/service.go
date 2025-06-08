package account

import (
	"context"

	"github.com/ishanwardhono/transfer-system/internal/entity/dto"
	"github.com/ishanwardhono/transfer-system/internal/repository/account"
)

type Service interface {
	Register(ctx context.Context, req dto.RegisterAccountRequest) (dto.RegisterAccountResponse, error)
}

type service struct {
	AccountRepo account.Repository
}

func NewService(accRepo account.Repository) Service {
	return &service{
		AccountRepo: accRepo,
	}
}
