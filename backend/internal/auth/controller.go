package auth

import "github.com/gofiber/fiber/v2"

type Controller interface {
	UserLogin(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
}