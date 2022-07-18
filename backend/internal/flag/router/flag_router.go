package router

import (
	"github.com/blastertwist/flag-dash/internal/flag"
	"github.com/blastertwist/flag-dash/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func InitializeFlagRouter(r fiber.Router, mw *middlewares.MiddlewareManager, fc flag.Controller){
	flagGroup := r.Group("/flag")

	flagGroup.Get("/", mw.UserAuthorized, fc.GetFlags)
	flagGroup.Get("/:id", mw.UserAuthorized, fc.GetFlag)
	flagGroup.Post("/new-flag", mw.UserAuthorized, fc.NewFlag)
	flagGroup.Patch("/:id", mw.UserAuthorized, fc.EditFlag)
	flagGroup.Delete("/:id", mw.UserAuthorized, fc.DeleteFlag)
}