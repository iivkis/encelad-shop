package services

import (
	"context"
	"encelad-shared/core/domain"
	"encelad-shared/core/ports"
)

type UserServiceImpl struct {
	repo ports.UserRepository
}

func NewUserService(
	repo ports.UserRepository,
) ports.UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) GetByID(ctx context.Context, id int64) (*domain.UserModel, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserServiceImpl) Create(ctx context.Context, firstname string, lastname string, password string) (*domain.UserModel, error) {
	pwdbytes, err := new(domain.UserModel).PasswordToHash(password)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(pwdbytes)
	return s.repo.Create(ctx, firstname, lastname, hashedPassword)
}

func (s *UserServiceImpl) Update(ctx context.Context, id int64, firstname string, lastname string) (*domain.UserModel, error) {
	return s.repo.Update(ctx, id, firstname, lastname)
}
