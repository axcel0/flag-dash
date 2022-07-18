package middlewares

import (
	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/pkg/logger"
)

type MiddlewareManager struct {
	cfg *config.Config
	logger logger.Logger
}

func NewMiddlewareManager(cfg *config.Config, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg:cfg, logger:logger}
}