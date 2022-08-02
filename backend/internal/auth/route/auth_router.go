package router

import (
	"github.com/blastertwist/flag-dash/internal/auth"
	"github.com/blastertwist/flag-dash/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func InitializeAuthRoute(r fiber.Router, mw *middlewares.MiddlewareManager, ac auth.Controller){
	
	authGroup := r.Group("/auth");

	authGroup.Post("/login", ac.UserLogin)
	authGroup.Get("/profile", mw.UserAuthorized, ac.GetUserProfile)
	authGroup.Post("/create-user", mw.UserAuthorized, ac.CreateUser)
	authGroup.Get("/", mw.UserAuthorized, ac.GetUser)
	authGroup.Post("/refresh-token", ac.RefreshToken)
}