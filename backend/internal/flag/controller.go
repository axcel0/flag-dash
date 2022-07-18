package flag

import "github.com/gofiber/fiber/v2"


type Controller interface {
	GetFlags(c *fiber.Ctx) error
	GetFlag(c * fiber.Ctx) error
	NewFlag(c *fiber.Ctx) error
	EditFlag(c *fiber.Ctx) error
	DeleteFlag(c *fiber.Ctx) error
}