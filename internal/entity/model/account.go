package model

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID        int64           `json:"id" db:"id"`
	Balance   decimal.Decimal `json:"balance" db:"balance"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at"`
	DeletedAt null.Time       `json:"deleted_at" db:"deleted_at"`
}
