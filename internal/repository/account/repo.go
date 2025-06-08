package account

import (
	"context"

	"github.com/ishanwardhono/transfer-system/internal/entity/model"
	"github.com/ishanwardhono/transfer-system/pkg/db"
)

type Repository interface {
	InsertAccount(ctx context.Context, account model.Account) error
}

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) InsertAccount(ctx context.Context, account model.Account) error {
	_, err := r.db.NamedExecContext(
		ctx,
		insertAccountQuery,
		account,
	)
	return err
}
