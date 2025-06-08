package account

import (
	"context"

	"github.com/ishanwardhono/transfer-system/internal/entity/dto"
	"github.com/ishanwardhono/transfer-system/internal/repository/account"
)

type Service interface {
	Register(ctx context.Context, req dto.RegisterAccountRequest) error
	GetById(ctx context.Context, accountId int64) (dto.GetAccountByIdResponse, error)
}

type service struct {
	accountRepo account.Repository
}

func NewService(accRepo account.Repository) Service {
	return &service{
		accountRepo: accRepo,
	}
}
