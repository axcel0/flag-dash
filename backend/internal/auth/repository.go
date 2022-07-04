package auth

import (
	"context"

	"github.com/blastertwist/flag-dash/internal/dao"
)

type Repository interface {
	Create(ctx context.Context, user *dao.User, userProfile *dao.UserProfile) (*dao.User, *dao.UserProfile, error)
	Update(ctx context.Context, user *dao.User) (*dao.User, error)
	Delete(ctx context.Context, userID uint32) error
	FindByID(ctx context.Context, userID uint32) (*dao.User, error)
	FindByEmail(ctx context.Context, email string) (*dao.User, error)
}