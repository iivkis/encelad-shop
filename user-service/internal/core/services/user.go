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
) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) GetByID(ctx context.Context, id int64) (*domain.UserModel, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserServiceImpl) Create(ctx context.Context, in *ports.CreateUserIn) (*domain.UserModel, error) {
	return s.repo.Create(ctx, in)
}

func (s *UserServiceImpl) Update(ctx context.Context, id int64, in *ports.UpdateUserIn) (*domain.UserModel, error) {
	return s.repo.Update(ctx, id, in)
}
