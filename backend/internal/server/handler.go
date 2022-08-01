package server

import (
	authController "github.com/blastertwist/flag-dash/internal/auth/controller"
	authRepository "github.com/blastertwist/flag-dash/internal/auth/repository"
	authRouter "github.com/blastertwist/flag-dash/internal/auth/route"
	authService "github.com/blastertwist/flag-dash/internal/auth/service"
	flagController "github.com/blastertwist/flag-dash/internal/flag/controller"
	flagRepository "github.com/blastertwist/flag-dash/internal/flag/repository"
	flagRouter "github.com/blastertwist/flag-dash/internal/flag/router"
	flagService "github.com/blastertwist/flag-dash/internal/flag/service"
	"github.com/blastertwist/flag-dash/internal/middlewares"
	projectController "github.com/blastertwist/flag-dash/internal/project/controller"
	projectRepository "github.com/blastertwist/flag-dash/internal/project/repository"
	projectRouter "github.com/blastertwist/flag-dash/internal/project/route"
	projectService "github.com/blastertwist/flag-dash/internal/project/service"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *Server) InitializeServer () {
	// Initialize Repositories
	authRepository := authRepository.NewAuthRepository(s.cfg, s.db, s.logger)
	projectRepository := projectRepository.NewProjectRepository(s.cfg, s.db, s.logger)
	flagRepository := flagRepository.NewFlagRepo(s.cfg, s.db)

	// Initialize Services
	authService := authService.NewAuthService(s.cfg, authRepository)
	projectService := projectService.NewProjectService(s.cfg, projectRepository)
	flagService := flagService.NewFlagService(s.cfg, flagRepository)

	// Initialize Handlers
	authController := authController.NewAuthController(s.cfg, authService)
	projectController := projectController.NewProjectController(s.cfg, projectService)
	flagController := flagController.NewFlagController(s.cfg, flagService)


	//Initialize Middlewares
	mw := middlewares.NewMiddlewareManager(s.cfg, s.logger)

	s.fiber.Use(compress.New(compress.Config{
    	Level: compress.LevelBestSpeed, // 1
	}))

	s.fiber.Use(cors.New())

	// Initialize Routes
	api := s.fiber.Group("/api")

	// Initialize API Version
	v1 := api.Group("/v1")

	// Initialize All Routes
	authRouter.InitializeAuthRoute(v1, mw, authController)
	projectRouter.InitializeProjectRoute(v1, mw, projectController)
	flagRouter.InitializeFlagRouter(v1, mw, flagController)

}