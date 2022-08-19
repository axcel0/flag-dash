package flag

import (
	"context"

	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/pkg/utils"
)

type Repository interface{
	GetAllFlags(ctx context.Context, projectID uint32) ([]*dao.Flag, error)

	GetFlagsCount(ctx context.Context, projectID uint32) (uint32, error)
	GetFlags(ctx context.Context, pq *utils.PaginationQuery, projectID uint32) ([]*dao.Flag, error)
	GetFlag(ctx context.Context, flag *dao.Flag) (*dao.Flag, error)
	NewFlag(ctx context.Context, flag *dao.Flag) (*dao.Flag, error)
	EditFlag(ctx context.Context, flag *dao.Flag) (*dao.Flag, error)
	DeleteFlag(ctx context.Context, flag *dao.Flag) error

	GetFlagContextsCount(ctx context.Context, projectID uint32) (uint32, error)
	GetFlagContexts(ctx context.Context, pq *utils.PaginationQuery, flagID uint32) ([]*dao.FlagContext, error)
	GetFlagContext(ctx context.Context, flagContext *dao.FlagContext) (*dao.FlagContext, error)
	NewFlagContext(ctx context.Context, flagContext *dao.FlagContext) (*dao.FlagContext, error)
	EditFlagContext(ctx context.Context, flagContext *dao.FlagContext) (*dao.FlagContext, error)
	DeleteFlagContext(ctx context.Context, flagContext *dao.FlagContext) error
}