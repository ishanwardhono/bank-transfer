package transaction

import (
	"context"

	"github.com/ishanwardhono/transfer-system/internal/entity/dto"
	"github.com/ishanwardhono/transfer-system/internal/repository/account"
	"github.com/ishanwardhono/transfer-system/internal/repository/dbtrx"
	"github.com/ishanwardhono/transfer-system/internal/repository/transaction"
)

type Service interface {
	Transfer(ctx context.Context, req dto.TransferRequest) error
}

type service struct {
	dbTrx           dbtrx.Repository
	transactionRepo transaction.Repository
	accountRepo     account.Repository
}

func NewService(
	dbTrx dbtrx.Repository,
	transRepo transaction.Repository,
	accountRepo account.Repository,
) Service {
	return &service{
		dbTrx:           dbTrx,
		transactionRepo: transRepo,
		accountRepo:     accountRepo,
	}
}
