package middlewares

import (
	"strconv"

	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/internal/dto"
	"github.com/blastertwist/flag-dash/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (mw *MiddlewareManager) ProjectAuthorized(c *fiber.Ctx) error {
	token := c.Get("ProjectAuthorization")

	if token == ""{
		return c.Status(fiber.ErrUnauthorized.Code).JSON(&dto.ProjectAuthorizationResponse{
			Status: "401",
			Msg: "Unauthorized Project, There is no token",
		})
	}

	isValid, claims, err := utils.VerifyJWTProject(token, mw.cfg.JWT.SecretKey)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	if !isValid {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(&dto.ProjectAuthorizationResponse{
			Status: "401",
			Msg: "Unauthorized Project, token is invalid.",
		})
	}

	iProjID, err := strconv.ParseUint(claims.ID, 32, 32)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	p, err := mw.pr.GetProjectAccessKey(c.Context(), &dao.Project{ID: uint32(iProjID)})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	if p.AccessKey != token {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(&dto.ProjectAuthorizationResponse{
			Status: "401",
			Msg: "Unauthorized Project, token provided not same as in the database.",
		})
	}

	return c.Next()
}