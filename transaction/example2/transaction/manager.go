package transaction

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type Manager interface {
	Run(
		ctx context.Context,
		callback func(ctx context.Context, tx *sql.Tx) error,
	) error
}

type txKey string

var ctxWithTx = txKey("tx")

type SQLTransactionManager struct {
	db *sql.DB
}

func NewManager(db *sql.DB) *SQLTransactionManager {
	return &SQLTransactionManager{db: db}
}

func (m *SQLTransactionManager) Run(
	ctx context.Context,
	callback func(ctx context.Context, tx *sql.Tx) error,
) (rErr error) {
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		if rErr != nil {
			rErr = multierr.Combine(rErr, errors.WithStack(tx.Rollback()))
		}
	}()

	defer func() {
		if rec := recover(); rec != nil {
			if e, ok := rec.(error); ok {
				rErr = e
			} else {
				rErr = errors.Errorf("%s", rec)
			}
		}
	}()

	if err = callback(ctx, tx); err != nil {
		return err
	}

	return errors.WithStack(tx.Commit())
}
