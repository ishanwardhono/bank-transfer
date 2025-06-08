package dto

import (
	"errors"

	"github.com/ishanwardhono/transfer-system/internal/entity/model"
	"github.com/shopspring/decimal"
)

type RegisterAccountRequest struct {
	AccountID      int64           `json:"account_id"`
	InitialBalance decimal.Decimal `json:"initial_balance"`
}

func (r *RegisterAccountRequest) Validate() error {
	if r.AccountID == 0 {
		return errors.New("id is required")
	}
	return nil
}

func (r RegisterAccountRequest) ToModel() model.Account {
	return model.Account{
		ID:      r.AccountID,
		Balance: r.InitialBalance,
	}
}
