package service

import (
	"brand/transaction/example2/model"
	"brand/transaction/example2/transaction"
	"context"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	WithTransaction(tx *sql.Tx) UserRepository
}

type ProfileRepository interface {
	Create(ctx context.Context, user *model.Profile) error
	WithTransaction(tx *sql.Tx) ProfileRepository
}

type RegistrationData struct {
	Email string
	Name  string
}

type RegistrationService struct {
	transactionManager transaction.Manager
	userRepository     UserRepository
	profileRepository  ProfileRepository
}

func NewRegistrationService(
	transactionManager transaction.Manager,
	userRepository UserRepository,
	profileRepository ProfileRepository,
) *RegistrationService {
	return &RegistrationService{
		transactionManager: transactionManager,
		userRepository:     userRepository,
		profileRepository:  profileRepository,
	}
}

func (s *RegistrationService) Register(ctx context.Context, data RegistrationData) error {
	return s.transactionManager.Run(ctx, func(ctx context.Context, tx *sql.Tx) error {
		userRepository := s.userRepository.WithTransaction(tx)
		profileRepository := s.profileRepository.WithTransaction(tx)

		if err := userRepository.Create(ctx, &model.User{
			Email: data.Email,
		}); err != nil {
			return err
		}

		if err := profileRepository.Create(ctx, &model.Profile{
			Name: data.Name,
		}); err != nil {
			return err
		}

		return nil
	})
}
