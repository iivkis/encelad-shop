package test

import (
	"context"
	"encelad-shared/domain"
	"encelad-shared/ports"
	"encelad-shared/ports/portsmock"

	"enceland_user-service/internal/core/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserServiceCreate(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := portsmock.NewMockUserRepository(ctrl)

	in := &ports.CreateUserIn{
		Firstname: "John",
		Lastname:  "Doe",
	}

	out := domain.NewUserModel(
		1,
		"John",
		"Doe",
		domain.UserRole().User,
	)

	repo.
		EXPECT().
		Create(ctx, in).
		Return(out, nil).
		Times(1)

	userService := services.NewUserService(repo)

	result, err := userService.Create(ctx, in)
	assert.Nil(t, err)
	assert.Equal(t, out, result)
}
