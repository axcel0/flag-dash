package controller

import (
	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/dto"
	"github.com/blastertwist/flag-dash/internal/flag"
	"github.com/gofiber/fiber/v2"
)

type flagController struct {
	cfg *config.Config
	fs	flag.Service
}

func NewFlagController(cfg *config.Config, fs flag.Service) flag.Controller{
	return &flagController{cfg:cfg, fs:fs}
}

func (fc *flagController) GetFlags(c *fiber.Ctx) error {
	userReq := &dto.GetFlagsRequest{}
	c.BodyParser(&userReq)

	res, err := fc.fs.GetFlags(c.Context(), userReq)
	if err != nil {
		return err
	}
	c.JSON(res)
	return nil
}

func (fc *flagController) GetFlag(c *fiber.Ctx) error {
	userReq := &dto.GetFlagRequest{}
	c.BodyParser(&userReq)

	res, err := fc.fs.GetFlag(c.Context(), userReq)
	if err != nil {
		return err
	}

	c.JSON(res)
	return nil
}

func (fc *flagController) NewFlag(c *fiber.Ctx) error {
	userReq := &dto.NewFlagRequest{}
	c.BodyParser(&userReq)

	res, err := fc.fs.NewFlag(c.Context(), userReq)
	if err != nil {
		return err
	}

	c.JSON(res)
	return nil
}

func (fc *flagController) EditFlag(c *fiber.Ctx) error {
	userReq := &dto.EditFlagRequest{}
	c.BodyParser(&userReq)

	res, err := fc.fs.EditFlag(c.Context(), userReq)
	if err != nil {
		return err
	}

	c.JSON(res)
	return nil
}

func (fc *flagController) DeleteFlag(c *fiber.Ctx) error {
	userReq := &dto.DeleteFlagRequest{}
	c.BodyParser(&userReq)

	res, err := fc.fs.DeleteFlag(c.Context(), userReq)
	if err != nil {
		return err
	}

	c.JSON(res)
	return nil
}