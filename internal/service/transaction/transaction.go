package transaction

import (
	"context"
	"net/http"

	"github.com/ishanwardhono/transfer-system/internal/entity/dto"
	"github.com/ishanwardhono/transfer-system/pkg/errors"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
	"github.com/ishanwardhono/transfer-system/pkg/utils"
)

func (s *service) Transfer(ctx context.Context, req dto.TransferRequest) error {
	sourceAccount, err := s.accountRepo.GetAccount(ctx, req.SourceAccountID)
	if err != nil {
		logger.Errorf(ctx, "failed to get source account, err: %v", err)
		return err
	}

	destinationAccount, err := s.accountRepo.GetAccount(ctx, req.DestinationAccountID)
	if err != nil {
		logger.Errorf(ctx, "failed to get destination account, err: %v", err)
		return err
	}

	sourceAccountBalance := sourceAccount.Balance.Sub(req.Amount)
	if sourceAccountBalance.IsNegative() {
		logger.Errorf(ctx, "insufficient balance in source account, balance: %s, amount: %s", sourceAccount.Balance.String(), req.Amount.String())
		return errors.New(http.StatusBadRequest, "insufficient balance in source account")
	}
	destinationAccountBalance := destinationAccount.Balance.Add(req.Amount)

	tx, err := s.dbTrx.BeginTx(ctx)
	if err != nil {
		logger.Errorf(ctx, "failed to begin transaction, err: %v", err)
		return err
	}

	transactionModel := req.ToModel()
	transactionModel.ReferenceNumber = utils.GenerateReferenceNumber()
	if err := s.transactionRepo.TxInsertTransaction(ctx, tx, transactionModel); err != nil {
		logger.Errorf(ctx, "failed to insert transaction, err: %v", err)
		s.dbTrx.RollbackTx(ctx, tx)
		return err
	}

	if err := s.accountRepo.TxUpdateBalance(ctx, tx, req.SourceAccountID, sourceAccountBalance); err != nil {
		logger.Errorf(ctx, "failed to update source account balance, err: %v", err)
		s.dbTrx.RollbackTx(ctx, tx)
		return err
	}

	if err := s.accountRepo.TxUpdateBalance(ctx, tx, req.DestinationAccountID, destinationAccountBalance); err != nil {
		logger.Errorf(ctx, "failed to update destination account balance, err: %v", err)
		s.dbTrx.RollbackTx(ctx, tx)
		return err
	}

	if err := tx.Commit(); err != nil {
		logger.Errorf(ctx, "failed to commit transaction, err: %v", err)
		s.dbTrx.RollbackTx(ctx, tx)
		return err
	}

	return nil
}
