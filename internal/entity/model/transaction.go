package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID                   int64           `json:"id" db:"id"`
	SourceAccountID      int64           `json:"source_account_id" db:"source_account_id"`
	DestinationAccountID int64           `json:"destination_account_id" db:"destination_account_id"`
	Amount               decimal.Decimal `json:"amount" db:"amount"`
	ReferenceNumber      string          `json:"reference_number" db:"reference_number"`
	CreatedAt            time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at" db:"updated_at"`
}
