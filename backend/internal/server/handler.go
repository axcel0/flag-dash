package server

import (
	"github.com/blastertwist/flag-dash/internal/auth/controller"
	"github.com/blastertwist/flag-dash/internal/auth/repository"
	router "github.com/blastertwist/flag-dash/internal/auth/route"
	"github.com/blastertwist/flag-dash/internal/auth/service"
)

func (s *Server) InitializeServer () {
	// Initialize Repositories
	authRepository := repository.NewAuthRepository(s.cfg, s.db, s.logger)

	// Initialize Services
	authService := service.NewAuthService(s.cfg, authRepository)

	// Initialize Handlers
	authController := controller.NewAuthController(s.cfg, authService)

	// Initialize Routes
	api := s.fiber.Group("/api")

	// Initialize API Version
	v1 := api.Group("/v1")

	// Initialize All Routes
	router.InitializeAuthRoute(v1, authController)

}