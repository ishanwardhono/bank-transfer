package transaction

import (
	"context"

	"github.com/ishanwardhono/transfer-system/internal/entity/model"
	"github.com/ishanwardhono/transfer-system/pkg/db"
	"github.com/jmoiron/sqlx"
)

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
	var id int64
	rows, err := r.db.NamedQueryContext(
		ctx,
		insertTransactionQuery,
		transaction,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return err
		}
	}

	return nil
}
