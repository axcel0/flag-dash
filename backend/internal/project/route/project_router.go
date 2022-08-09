package router

import (
	"github.com/blastertwist/flag-dash/internal/middlewares"
	"github.com/blastertwist/flag-dash/internal/project"
	"github.com/gofiber/fiber/v2"
)

func InitializeProjectRoute(r fiber.Router, mw *middlewares.MiddlewareManager, pc project.Controller){
	projectGroup := r.Group("/project")

	projectGroup.Get("/", mw.UserAuthorized, pc.GetProjects)
	projectGroup.Get("/:id", pc.GetProjectByID)

	projectGroup.Post("/new-project", mw.UserAuthorized, pc.NewProject)
	projectGroup.Patch("/:id", mw.UserAuthorized, pc.EditProject)
	projectGroup.Delete("/:id",  mw.UserAuthorized, pc.DeleteProject)
}