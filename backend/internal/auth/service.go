package auth

import (
	"context"

	"github.com/blastertwist/flag-dash/internal/dto"
)

type Service interface {
	GetUsers(ctx context.Context, getUsersReq *dto.GetUsersRequest) (*dto.GetUsersResponse, error)
	UserLogin(ctx context.Context, cu *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
	CreateUser(ctx context.Context, cu *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	GetUserByEmail(ctx context.Context, gu *dto.GetUserRequest) (*dto.GetUserResponse, error)
	GetUserByID(ctx context.Context, userID uint32) (*dto.GetUserProfileResponse, error)
	EditUser(ctx context.Context, editUserReq *dto.EditUserRequest) (*dto.EditUserResponse, error)
	DeleteUser(ctx context.Context, delUserReq *dto.DeleteUserRequest) (*dto.DeleteUserResponse, error)
	RefreshToken(ctx context.Context, refreshTokenReq *dto.UserRefreshTokenRequest) (*dto.UserRefreshTokenResponse, error)
}
