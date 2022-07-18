package project

import (
	"context"

	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/pkg/utils"
)

type Repository interface {
	NewProject(ctx context.Context, project *dao.Project)(*dao.Project, error)
	EditProject(ctx context.Context, project *dao.Project) (*dao.Project, error)
	DeleteProject(ctx context.Context, project *dao.Project) error
	GetProject(ctx context.Context, project *dao.Project) (*dao.Project, error)
	GetProjects(ctx context.Context, pq *utils.PaginationQuery) ([]*dao.Project, error)
	GetProjectCount(ctx context.Context)(uint8, error)
}