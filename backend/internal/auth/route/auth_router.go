package router

import (
	"github.com/blastertwist/flag-dash/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func InitializeAuthRoute(r fiber.Router, ac auth.Controller){
	
	authGroup := r.Group("/auth");

	authGroup.Post("/login", ac.UserLogin)
	authGroup.Post("/create-user", ac.CreateUser)
	authGroup.Get("/", ac.GetUser)
}