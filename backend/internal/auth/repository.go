package auth

import (
	"context"

	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/pkg/utils"
)

//go:generate mockgen -source repository.go -destination mock/repository_mock.go -package mock
type Repository interface {
	GetUsersCount(ctx context.Context) (uint8, error)
	GetUsers(ctx context.Context, pq *utils.PaginationQuery) ([]*dao.User, error)
	CreateUser(ctx context.Context, user *dao.User, userProfile *dao.UserProfile) (*dao.User, *dao.UserProfile, error)
	Update(ctx context.Context, user *dao.User) (*dao.User, error)
	Delete(ctx context.Context, user *dao.User) error
	FindByID(ctx context.Context, user *dao.User) (*dao.User, error)
	FindByEmail(ctx context.Context, user *dao.User) (*dao.User, error)
}
