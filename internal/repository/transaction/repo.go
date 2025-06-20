package transaction

import (
	"context"

	"github.com/ishanwardhono/transfer-system/internal/entity/model"
	"github.com/ishanwardhono/transfer-system/pkg/db"
	"github.com/jmoiron/sqlx"
)

//go:generate go run go.uber.org/mock/mockgen --source=repo.go --package=mocktransactionrepo --destination=../../../test/mock/repository/transaction/repo.go
type Repository interface {
	TxInsertTransaction(ctx context.Context, tx *sqlx.Tx, transaction model.Transaction) error
}

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) TxInsertTransaction(ctx context.Context, tx *sqlx.Tx, transaction model.Transaction) error {
	_, err := tx.NamedExecContext(
		ctx,
		insertTransactionQuery,
		transaction,
	)
	if err != nil {
		return err
	}

	return nil
}
