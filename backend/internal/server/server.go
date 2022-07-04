package server

import (
	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	fiber *fiber.App
	cfg *config.Config
	db *sqlx.DB
	logger logger.Logger
}

func NewServer(cfg *config.Config, db *sqlx.DB, logger logger.Logger) *Server{
	return &Server{
		fiber:fiber.New(),
		cfg:cfg,
		db:db,
		logger:logger,
	}
}

func (s *Server) Run (){
	s.InitializeServer()

	s.fiber.Listen(":3001")

}