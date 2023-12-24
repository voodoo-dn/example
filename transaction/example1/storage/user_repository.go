package storage

import (
	"brand/transaction"
	"brand/transaction/example1/model"
	"context"
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
