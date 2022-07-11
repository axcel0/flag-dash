package auth

import (
	"context"

	"github.com/blastertwist/flag-dash/internal/dao"
)

//go:generate mockgen -source repository.go -destination mock/repository_mock.go -package mock
type Repository interface {
	CreateUser(ctx context.Context, user *dao.User, userProfile *dao.UserProfile) (*dao.User, *dao.UserProfile, error)
	Update(ctx context.Context, user *dao.User) (*dao.User, error)
	Delete(ctx context.Context, user *dao.User) error
	FindByID(ctx context.Context, user *dao.User) (*dao.User, error)
	FindByEmail(ctx context.Context, user *dao.User) (*dao.User, error)
}
