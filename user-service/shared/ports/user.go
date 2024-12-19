package ports

//go:generate mockgen -destination=./portsmock/user.go -package=portsmock . UserRepository,UserService

import (
	"context"
	"encelad-shared/domain"
)

type CreateUserIn struct {
	Firstname string
	Lastname  string
}

type UpdateUserIn struct {
	Firstname string
	Lastname  string
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.UserModel, error)
	Create(ctx context.Context, in *CreateUserIn) (*domain.UserModel, error)
	Update(ctx context.Context, id int64, in *UpdateUserIn) (*domain.UserModel, error)
}

type UserService interface {
	GetByID(ctx context.Context, id int64) (*domain.UserModel, error)
	Create(ctx context.Context, in *CreateUserIn) (*domain.UserModel, error)
	Update(ctx context.Context, id int64, in *UpdateUserIn) (*domain.UserModel, error)
}
