package storage

import (
	"brand/transaction"
	"brand/transaction/example1/model"
	"context"
)

type ProfileRepository struct {
	db transaction.DB
}

func NewProfileRepository(db transaction.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) Create(ctx context.Context, profile *model.Profile) error {
	_, err := r.db.ExecContext(ctx, "INSERT ...", profile.Name)

	return err
}
