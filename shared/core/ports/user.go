package ports

//go:generate mockgen -destination=./portsmock/user.go -package=portsmock . UserRepository,UserService

import (
	"context"
	"encelad-shared/core/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.UserModel, error)
	Create(ctx context.Context, firstname string, lastname string, hashedPassword string) (*domain.UserModel, error)
	Update(ctx context.Context, id int64, firstname string, lastname string) (*domain.UserModel, error)
}

type UserService interface {
	GetByID(ctx context.Context, id int64) (*domain.UserModel, error)
	Create(ctx context.Context, firstname string, lastname string, password string) (*domain.UserModel, error)
	Update(ctx context.Context, id int64, firstsname string, lastname string) (*domain.UserModel, error)
}
