package controller

import (
	"database/sql"
	"fmt"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/dto"
	"github.com/blastertwist/flag-dash/internal/project"
	"github.com/gofiber/fiber/v2"
)

type projectController struct {
	cfg *config.Config
	ps project.Service
}

func NewProjectController(cfg *config.Config, ps project.Service) project.Controller{
	return &projectController{cfg:cfg, ps:ps}
}

func (pc *projectController) GetProjects (c *fiber.Ctx) error {
	getProjectsReq := &dto.GetProjectsRequest{}
	if err := c.QueryParser(getProjectsReq); err != nil {
            return err
        }
	res, err := pc.ps.GetProjects(c.Context(), getProjectsReq)
	
	if err != nil {
		fmt.Print(err)
		return err
	}

	return c.JSON(res)
}

func (pc *projectController) GetProjectByID (c *fiber.Ctx) error {
	getProjectReq := &dto.GetProjectRequest{}
	c.ParamsParser(getProjectReq);

	res, err := pc.ps.GetProject(c.Context(), getProjectReq)

	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (pc *projectController) NewProject (c *fiber.Ctx) error {
	newProjectReq := &dto.NewProjectRequest{}
	c.BodyParser(newProjectReq)

	res, err := pc.ps.NewProject(c.Context(), newProjectReq)

	if err != nil {
		return err
	}

	c.JSON(res)
	return nil
}

func (pc *projectController) EditProject (c *fiber.Ctx) error {
	editProjectReq := &dto.EditProjectRequest{}
	c.ParamsParser(editProjectReq);
	c.BodyParser(editProjectReq)

	projectID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	editProjectReq.ID = uint32(projectID)

	res, err := pc.ps.EditProject(c.Context(), editProjectReq)
	
	if err != nil {
		return err
	}

	c.JSON(res)
	return nil
}

func (pc *projectController) DeleteProject (c *fiber.Ctx) error {
	deleteProjectReq := &dto.DeleteProjectRequest{}
	c.ParamsParser(deleteProjectReq)

	err := pc.ps.DeleteProject(c.Context(), deleteProjectReq)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return c.Status(fiber.StatusNotFound).JSON("Project Not Found.");
		} else {
			return err
		}
	}

	return 	c.Status(fiber.StatusOK).JSON(&dto.DeleteProjectResponse{
		Status: "200",
		Msg: "Delete Project Success",
	})
}

func (pc *projectController) GenerateProjectAccessKey(c *fiber.Ctx) error {
	genProjectAccessKeyReq := &dto.GenerateProjectAccessKeyRequest{}
	c.ParamsParser(genProjectAccessKeyReq)

	res, err := pc.ps.GenerateProjectAccessKey(c.Context(), genProjectAccessKeyReq)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}