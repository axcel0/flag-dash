package service

import (
	"context"
	"log"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/auth"
	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/internal/dto"
)

type authService struct {
	cfg 	*config.Config
	r 		auth.Repository
}

func NewAuthService(cfg *config.Config, r auth.Repository) auth.Service {
	return &authService{cfg:cfg, r:r}
}

func(s *authService) UserLogin(cu *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	return &dto.UserLoginResponse{}, nil
}

func(s *authService) CreateUser(cu *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	_, _, err := s.r.Create(context.Background(), &dao.User{
		Email: cu.Email,
		Password: cu.Password,
	},
	&dao.UserProfile{
		FirstName: cu.FirstName,
		LastName: cu.LastName,
		PhoneNumber: cu.PhoneNumber,
	})

	if err != nil {
		log.Fatal(err)
		return &dto.CreateUserResponse{
			Status: "401",
			Msg: "Failed to create user",
		}, err
	}
	return &dto.CreateUserResponse{
		Status: "201",
		Msg: "User created.",
	}, nil
}

func(s *authService) GetUserByEmail(gu *dto.GetUserRequest) (*dto.GetUserResponse, error){
	u, err := s.r.FindByEmail(context.Background(), gu.Email)
	if err != nil {
		return &dto.GetUserResponse{
			Status: "500",
			Msg: "Internal Server Error, failed to get User By Email.",
		}, err
	}

	userRes := &dto.GetUserResponse{
		Status: "200",
		Msg: "Successfully get user by e-mail",
	}
	userRes.User.Email = u.Email
	userRes.User.Profile.FirstName = u.FirstName
	userRes.User.Profile.LastName = u.LastName
	userRes.User.Role.Name = u.RoleName
	userRes.User.Role.Level = u.RoleLevel

	return userRes, nil
}
