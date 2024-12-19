package test

import (
	"context"
	"encelad-shared/core/domain"
	"encelad-shared/core/ports/portsmock"
	"enceland_user-service/internal/core/services"
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserServiceCreate(t *testing.T) {
	type tableItem struct {
		firstname string
		lastname  string
		password  string
	}

	table := []*tableItem{
		{
			firstname: "John",
			lastname:  "Doe",
			password:  "password",
		},
	}

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := portsmock.NewMockUserRepository(ctrl)

	repo.
		EXPECT().
		Create(
			ctx,
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).
		DoAndReturn(func(ctx context.Context, firstname string, lastname string, hashedPassword string) (*domain.UserModel, error) {
			return domain.NewUserModel(
				1,
				firstname,
				lastname,
				domain.UserRole().User,
				hashedPassword,
			), nil
		})

	service := services.NewUserService(repo)

	for _, tt := range table {
		user, err := service.Create(
			ctx,
			tt.firstname,
			tt.lastname,
			tt.password,
		)
		fmt.Println(user)
		assert.Nil(t, err)
		assert.Equal(t, tt.firstname, user.GetFirstname())
		assert.Equal(t, tt.lastname, user.GetLastname())
		assert.Equal(t, domain.UserRole().User, user.GetRole())
		assert.True(t, user.ComparePassword(tt.password))
	}
}
