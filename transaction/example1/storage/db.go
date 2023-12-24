package storage

import (
	"brand/transaction/example1/transaction"
	"context"
	"database/sql"
)

type DB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{db: db}
}

func (d *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	tx, ok := transaction.ExtractTxFromContext(ctx)
	if !ok {
		return d.db.QueryRowContext(ctx, query, args...)
	}

	return tx.QueryRowContext(ctx, query, args...)
}

func (d *DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	tx, ok := transaction.ExtractTxFromContext(ctx)
	if !ok {
		return d.db.QueryContext(ctx, query, args...)
	}

	return tx.QueryContext(ctx, query, args...)
}

func (d *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	tx, ok := transaction.ExtractTxFromContext(ctx)
	if !ok {
		return d.db.ExecContext(ctx, query, args...)
	}

	return tx.ExecContext(ctx, query, args...)
}

func (d *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	tx, ok := transaction.ExtractTxFromContext(ctx)
	if !ok {
		return d.db.PrepareContext(ctx, query)
	}

	return tx.PrepareContext(ctx, query)
}
