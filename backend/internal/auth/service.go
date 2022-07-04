package auth

import (
	"github.com/blastertwist/flag-dash/internal/dto"
)

type Service interface {
	UserLogin(cu *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
	CreateUser(cu *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	GetUserByEmail(gu *dto.GetUserRequest) (*dto.GetUserResponse, error)
}