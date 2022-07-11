package auth

import (
	"context"

	"github.com/blastertwist/flag-dash/internal/dto"
)

type Service interface {
	UserLogin(ctx context.Context, cu *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
	CreateUser(ctx context.Context, cu *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	GetUserByEmail(ctx context.Context, gu *dto.GetUserRequest) (*dto.GetUserResponse, error)
}
