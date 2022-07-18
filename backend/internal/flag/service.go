package flag

import (
	"context"

	"github.com/blastertwist/flag-dash/internal/dto"
)

type Service interface {
	GetFlags(ctx context.Context, getFlagsReq *dto.GetFlagsRequest) (*dto.GetFlagsResponse, error)
	GetFlag(ctx context.Context, getFlagReq *dto.GetFlagRequest) (*dto.GetFlagResponse, error)
	NewFlag(ctx context.Context, newFlagReq *dto.NewFlagRequest) (*dto.NewFlagResponse, error)
	EditFlag(ctx context.Context, editFlagReq *dto.EditFlagRequest) (*dto.EditFlagResponse, error)
	DeleteFlag(ctx context.Context, deleteFlagReq *dto.DeleteFlagRequest) (*dto.DeleteFlagResponse, error)
	GetFlagContexts(ctx context.Context, getFlagContextsReq *dto.GetFlagContextsRequest) (*dto.GetFlagContextsResponse, error)
	GetFlagContext(ctx context.Context, getFlagContextReq *dto.GetFlagContextRequest) (*dto.GetFlagContextResponse, error)
	NewFlagContext(ctx context.Context, newFlagContextReq *dto.NewFlagContextRequest) (*dto.NewFlagContextResponse, error)
	EditFlagContext(ctx context.Context, editFlagContextReq *dto.EditFlagContextRequest) (*dto.EditFlagContextResponse, error)
	DeleteFlagContext(ctx context.Context, deleteFlagContextReq *dto.DeleteFlagContextRequest) (*dto.DeleteFlagContextResponse, error)
}