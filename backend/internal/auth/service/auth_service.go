package service

import (
	"context"
	"errors"
	"log"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/auth"
	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/internal/dto"
	"github.com/blastertwist/flag-dash/pkg/utils"
)

type authService struct {
	cfg *config.Config
	r   auth.Repository
}

func NewAuthService(cfg *config.Config, r auth.Repository) auth.Service {
	return &authService{cfg: cfg, r: r}
}

func (s *authService) GetUsers(ctx context.Context, getUsersReq *dto.GetUsersRequest) (*dto.GetUsersResponse, error) {
	num, err := s.r.GetUsersCount(ctx)

	if err != nil {
		return nil, err
	}
	pq, err := utils.NewPagination(getUsersReq.Filter, getUsersReq.Limit, getUsersReq.PageNum, float32(num))

	if err != nil {
		return nil, err
	}

	users, err := s.r.GetUsers(ctx, pq)

	if err != nil {
		return nil, err
	}

	res := &dto.GetUsersResponse{
		Limit: pq.Limit,
		PageNum: pq.PageNum,
		MaxPage: pq.MaxNum,
	}

	for i := 0; i < len(users); i++ {
		ur := &dto.UserResponse{}
		ur.ID = users[i].ID
		ur.Email = users[i].Email
		ur.Profile.FirstName = users[i].FirstName
		ur.Profile.LastName = users[i].LastName
		ur.Role.Level = users[i].RoleLevel
		ur.Role.Name = users[i].RoleName

		res.Users = append(res.Users, ur)
	}

	return res, nil
}

func (s *authService) UserLogin(ctx context.Context, cu *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {

	user, errRepo := s.r.FindByEmail(ctx, &dao.User{Email: cu.Email})

	if errRepo != nil {
		return nil, errRepo
	}
	
	valid, errValid := utils.ValidatePassword(cu.Password, user.Password)

	if errValid != nil { 
		return nil, errValid
	}

	if !valid {
		return nil, errors.New("Password not same")
	}

	tokenNormal, errNormal := utils.GenerateJWT(user, s.cfg.JWT.SecretKey, 1)

	if errNormal != nil {
		return nil, errNormal
	}

	tokenRefresh, errRefresh := utils.GenerateJWT(user, s.cfg.JWT.RefreshSecretKey, 60)

	if errRefresh != nil {
		return nil, errRefresh
	}


	return &dto.UserLoginResponse{
		NormalJWT: tokenNormal,
		RefreshJWT: tokenRefresh,
	}, nil
}

func (s *authService) RefreshToken(ctx context.Context, refreshTokenReq *dto.UserRefreshTokenRequest) (*dto.UserRefreshTokenResponse, error) {
	isValid, claims, err := utils.VerifyJWT(refreshTokenReq.RefreshToken, s.cfg.JWT.RefreshSecretKey)

	if err != nil {
		return nil, err
	}

	if !isValid {
		return nil, errors.New("JWT Error: InvalidToken")
	}

	user, err := s.r.FindByEmail(ctx, &dao.User{Email: claims.Email})

	if err != nil {
		return nil, err
	}

	tokenString, err := utils.GenerateJWT(user, s.cfg.JWT.SecretKey, 1)

	if err != nil {
		return nil, err
	}


	return &dto.UserRefreshTokenResponse{
		Status: "200",
		Token: tokenString,
	}, nil
}

func (s *authService) CreateUser(ctx context.Context, cu *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	// Hashing Password
	hashedPassword, errHash := utils.HashPassword(cu.Password)
	if errHash != nil {
		return &dto.CreateUserResponse{
			Status: "401",
			Msg:    "Failed to create user, Hashing Failed",
		}, errHash
	}
	_, _, err := s.r.CreateUser(ctx, &dao.User{
		Email:    cu.Email,
		Password: hashedPassword,
	},
		&dao.UserProfile{
			FirstName:   cu.FirstName,
			LastName:    cu.LastName,
			PhoneNumber: cu.PhoneNumber,
		})

	if err != nil {
		log.Fatal(err)
		return &dto.CreateUserResponse{
			Status: "401",
			Msg:    "Failed to create user",
		}, err
	}
	return &dto.CreateUserResponse{
		Status: "201",
		Msg:    "User created.",
	}, nil
}

func (s *authService) EditUser(ctx context.Context, editUserReq *dto.EditUserRequest) (*dto.EditUserResponse, error) {
	
	u := &dao.User{
		ID: editUserReq.UserID,
		Email: editUserReq.Email,
		FirstName: editUserReq.FirstName,
		LastName: editUserReq.LastName,
		PhoneNumber: editUserReq.PhoneNumber,
		RoleLevel: editUserReq.RoleLevel,
	}

	u, err := s.r.Update(ctx, u)

	if err != nil {
		return nil, err
	}
	userRes := &dto.EditUserResponse{
		Status: "200",
		Msg:    "Successfully edit user.",
	}
	userRes.User.Email = u.Email
	userRes.User.Profile.FirstName = u.FirstName
	userRes.User.Profile.LastName = u.LastName
	userRes.User.Profile.PhoneNumber = u.PhoneNumber
	userRes.User.Role.Name = u.RoleName
	userRes.User.Role.Level = u.RoleLevel

	return userRes, nil
}

func (s *authService) GetUserByEmail(ctx context.Context, gu *dto.GetUserRequest) (*dto.GetUserResponse, error) {
	u, err := s.r.FindByEmail(ctx, &dao.User{
		Email: gu.Email,
	})
	if err != nil {
		return &dto.GetUserResponse{
			Status: "500",
			Msg:    "Internal Server Error, failed to get User By Email.",
		}, err
	}

	userRes := &dto.GetUserResponse{
		Status: "200",
		Msg:    "Successfully get user by e-mail",
	}
	userRes.User.Email = u.Email
	userRes.User.Profile.FirstName = u.FirstName
	userRes.User.Profile.LastName = u.LastName
	userRes.User.Role.Name = u.RoleName
	userRes.User.Role.Level = u.RoleLevel

	return userRes, nil
}

func (s *authService) GetUserByID(ctx context.Context, userID uint32) (*dto.GetUserProfileResponse, error) {
	u, err := s.r.FindByID(ctx, &dao.User{ID: userID})

	if err != nil {
		return nil, err
	}

	userRes := &dto.GetUserProfileResponse{
		Status: "200",
	}
	
	userRes.User.ID = u.ID
	userRes.User.Email = u.Email
	userRes.User.Profile.FirstName = u.FirstName
	userRes.User.Profile.LastName = u.LastName
	userRes.User.Role.Name = u.RoleName
	userRes.User.Role.Level = u.RoleLevel

	return userRes, nil
}

func (s *authService) DeleteUser(ctx context.Context, deleteUserReq *dto.DeleteUserRequest) (*dto.DeleteUserResponse, error) {
	err := s.r.Delete(ctx, &dao.User{ID: deleteUserReq.UserID})

	if err != nil {
		return nil, err
	}

	return &dto.DeleteUserResponse{
		Status: "200",
		Msg: "Delete user success",
	}, nil
}
