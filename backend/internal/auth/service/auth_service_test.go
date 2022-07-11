package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/auth/mock"
	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/internal/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAuthSC_FindUserByEmail(t *testing.T) {
	t.Parallel()

	cfg := &config.Config{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mock.NewMockRepository(ctrl)
	asc := NewAuthService(cfg, mr)

	ctx := context.Background()

	user := &dao.User{
		Email: "test@example.com",
	}

	mr.EXPECT().FindByEmail(ctx, gomock.Eq(user)).Return(nil, sql.ErrNoRows)

	_, err := asc.GetUserByEmail(ctx, &dto.GetUserRequest{Email: user.Email})

	require.Error(t, err)

}

func TestAuthSC_CreateUser(t *testing.T) {
	cfg := &config.Config{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mock.NewMockRepository(ctrl)
	asc := NewAuthService(cfg, mr)

	ctx := context.Background()

	user := &dao.User{
		Email:     "test@example.com",
		Password:  "test123",
		FirstName: "John",
		LastName:  "Doe",
	}

	userProfile := &dao.UserProfile{
		UserID:    1,
		FirstName: "John",
		LastName:  "Doe",
	}

	createUserReq := &dto.CreateUserRequest{
		Email:     "test@example.com",
		Password:  "test123",
		FirstName: "John",
		LastName:  "Doe",
	}

	// createUserRes := &dto.CreateUserResponse{
	// 	Status: "201",
	// 	Msg:    "User created.",
	// }

	//mr.EXPECT().FindByEmail(ctx, gomock.Eq(user)).Return(nil, sql.ErrNoRows)
	mr.EXPECT().CreateUser(ctx, gomock.Any(), gomock.Any()).Return(user, userProfile, nil)

	res, err := asc.CreateUser(ctx, createUserReq)

	require.NotNil(t, res)
	require.Nil(t, err)
}
