package middlewares

import (
	"errors"
	"fmt"

	"github.com/blastertwist/flag-dash/internal/dto"
	"github.com/blastertwist/flag-dash/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (m *MiddlewareManager) UserAuthorized(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return 	c.JSON(&dto.UserAuthorizedResponse{
			Status: "401",
			Msg: "There is no token found, please try to re-login to proceed.",
		})
	}

	isValid, claims, err := utils.VerifyJWT(token, m.cfg.JWT.SecretKey)
	fmt.Print(err)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return c.Status(fiber.ErrUnauthorized.Code).JSON(&dto.UserAuthorizedResponse{
				Status: "401",
				Msg: "Your token is expired, please try to re-login to proceed.",
			})
		}
		return err
	}
	if isValid {
		c.Locals("user_id", claims.ID)
		c.Next()
	} else {
		c.JSON(&dto.UserAuthorizedResponse{
			Status: "401",
			Msg: "Your token is invalid, please try to re-login to proceed.",
		})
	}
	return nil
}