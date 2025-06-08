package dbtrx

import (
	"context"

	"github.com/ishanwardhono/transfer-system/pkg/db"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
	"github.com/jmoiron/sqlx"
)

//go:generate go run go.uber.org/mock/mockgen --source=repo.go --package=mockdbtrxrepo --destination=../../../test/mock/repository/dbtrx/repo.go
type Repository interface {
	BeginTx(ctx context.Context) (*sqlx.Tx, error)
	RollbackTx(ctx context.Context, tx *sqlx.Tx)
	CommitTx(tx *sqlx.Tx) error
}

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) BeginTx(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := r.db.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (r *repository) RollbackTx(ctx context.Context, tx *sqlx.Tx) {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		logger.Errorf(ctx, "failed to rollback transaction, err: %v", rollbackErr)
	}
	return
}

func (r *repository) CommitTx(tx *sqlx.Tx) error {
	return tx.Commit()
}
