package account

import (
	"context"
	"net/http"

	"github.com/ishanwardhono/transfer-system/internal/entity/model"
	"github.com/ishanwardhono/transfer-system/pkg/db"
	"github.com/ishanwardhono/transfer-system/pkg/errors"
	"github.com/lib/pq"
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
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // 23505 is the PostgreSQL error code for unique violation
			return errors.New(http.StatusBadRequest, "account already exists")
		}
		return err
	}
	return err
}
