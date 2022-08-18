package middlewares

import (
	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/project"
	"github.com/blastertwist/flag-dash/pkg/logger"
)

type MiddlewareManager struct {
	pr project.Repository
	cfg *config.Config
	logger logger.Logger
}

func NewMiddlewareManager(pr project.Repository, cfg *config.Config, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{pr:pr, cfg:cfg, logger:logger}
}