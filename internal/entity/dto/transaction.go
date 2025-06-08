package dto

import (
	"errors"

	"github.com/ishanwardhono/transfer-system/internal/entity/model"
	"github.com/shopspring/decimal"
)

type TransferRequest struct {
	SourceAccountID      int64           `json:"source_account_id" validate:"required"`
	DestinationAccountID int64           `json:"destination_account_id" validate:"required"`
	Amount               decimal.Decimal `json:"amount" validate:"required"`
}

func (r TransferRequest) Validate() error {
	if r.SourceAccountID <= 0 {
		return errors.New("invalid source account ID")
	}
	if r.DestinationAccountID <= 0 {
		return errors.New("invalid destination account ID")
	}
	if r.Amount.IsNegative() || r.Amount.IsZero() {
		return errors.New("amount must be greater than zero")
	}
	return nil
}

func (r TransferRequest) ToModel() model.Transaction {
	return model.Transaction{
		SourceAccountID:      r.SourceAccountID,
		DestinationAccountID: r.DestinationAccountID,
		Amount:               r.Amount,
	}
}
