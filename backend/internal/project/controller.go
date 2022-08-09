package project

import "github.com/gofiber/fiber/v2"

type Controller interface {
	GetProjectByID(c *fiber.Ctx) error
	NewProject(c *fiber.Ctx) error
	EditProject(c *fiber.Ctx) error
	DeleteProject(c *fiber.Ctx) error
	GetProjects(c *fiber.Ctx) error
}