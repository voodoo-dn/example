package service

import (
	"brand/transaction/example1/model"
	"brand/transaction/example1/transaction"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
}

type ProfileRepository interface {
	Create(ctx context.Context, user *model.Profile) error
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
	return s.transactionManager.Run(ctx, func(ctx context.Context) error {
		if err := s.userRepository.Create(ctx, &model.User{
			Email: data.Email,
		}); err != nil {
			return err
		}

		if err := s.profileRepository.Create(ctx, &model.Profile{
			Name: data.Name,
		}); err != nil {
			return err
		}

		return nil
	})
}
