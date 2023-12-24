package storage

import (
	"brand/transaction"
	"brand/transaction/example2/model"
	"brand/transaction/example2/service"
	"context"
	"database/sql"
)

type UserRepository struct {
	db transaction.DB
}

func NewUserRepository(db transaction.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT ...", user.Email)

	return err
}

func (r *UserRepository) WithTransaction(tx *sql.Tx) service.UserRepository {
	return NewUserRepository(tx)
}
