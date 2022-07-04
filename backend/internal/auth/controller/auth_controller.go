package controller

import (
	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/auth"
	"github.com/blastertwist/flag-dash/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	cfg *config.Config
	as auth.Service
}

func NewAuthController(cfg *config.Config, as auth.Service) auth.Controller{
	return &authController{cfg: cfg, as:as}
}

func (ac *authController) UserLogin(c *fiber.Ctx) error {
	return nil
}

func (ac *authController) CreateUser(c *fiber.Ctx) error {
	userReq := &dto.CreateUserRequest{}
	c.BodyParser(userReq)

	res, err := ac.as.CreateUser(userReq)
	if err != nil {
		c.JSON(res)
		return err
	}

	c.JSON(res)
	return nil
}

func (ac *authController) GetUser(c *fiber.Ctx) error {
	getUserReq := &dto.GetUserRequest{}
	c.BodyParser(getUserReq)

	getUserRes, err := ac.as.GetUserByEmail(getUserReq)

	if err != nil {
		c.JSON(getUserRes)
		return err
	}
	c.JSON(getUserRes)

	return nil
}