package controller

import (
	"errors"
	"strconv"
	"time"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/auth"
	"github.com/blastertwist/flag-dash/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	cfg *config.Config
	as  auth.Service
}

func NewAuthController(cfg *config.Config, as auth.Service) auth.Controller {
	return &authController{cfg: cfg, as: as}
}


// UserLogin Function to handler UserLogin request.
// @Description Handle User Login.
// @Summary UserLogin controller function
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} dto.UserLoginResponse
// @Router /api/v1/auth/login [post]
func (ac *authController) UserLogin(c *fiber.Ctx) error {
	
	req := &dto.UserLoginRequest{}
	c.BodyParser(req)
	
	res, err := ac.as.UserLogin(c.Context(), req)

	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(err)
	}

	expirationTime := time.Now().Add(time.Duration(2) * time.Minute)

	jwtCookie := &fiber.Cookie{
		Name: "refreshToken_cookie",
		Value: res.RefreshJWT,
		Secure: false,
		HTTPOnly: true,
		Expires: expirationTime,
	}
	c.Cookie(jwtCookie)

	return c.Status(fiber.StatusOK).JSON(res)
}

func (ac *authController) GetUserProfile(c *fiber.Ctx) error {
	inUserID := c.Locals("user_id")
	sUserID, ok := inUserID.(string)

	if !ok {
		return errors.New("User ID is not valid")
	}
	
	iUserID, err := strconv.ParseUint(sUserID, 32, 32)

	if err != nil {
		return err
	}


	res, err := ac.as.GetUserByID(c.Context(), uint32(iUserID))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// CreateUser Function to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} dto.CreateUserResponse
// @Router /api/v1/auth/create-user [post]
func (ac *authController) CreateUser(c *fiber.Ctx) error {
	userReq := &dto.CreateUserRequest{}
	c.BodyParser(userReq)

	res, err := ac.as.CreateUser(c.Context(), userReq)
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

	getUserRes, err := ac.as.GetUserByEmail(c.Context(), getUserReq)

	if err != nil {
		c.JSON(getUserRes)
		return err
	}
	c.JSON(getUserRes)

	return nil
}

func (ac *authController) EditUser(c *fiber.Ctx) error {
	editUserReq  := &dto.EditUserRequest{}
	c.BodyParser(editUserReq)

	res, err := ac.as.EditUser(c.Context(), editUserReq)

	if err != nil {
		return err
	}

	c.JSON(res)
	return nil
}

func (ac *authController) DeleteUser (c *fiber.Ctx) error {
	deleteUserReq := &dto.DeleteUserRequest{}
	c.BodyParser(deleteUserReq)

	res, err := ac.as.DeleteUser(c.Context(), deleteUserReq)

	if err != nil {
		return err
	}

	c.JSON(res)
	return nil
}

func (ac *authController) RefreshToken(c *fiber.Ctx) error {
	refreshTokenReq := &dto.UserRefreshTokenRequest{}
	refreshTokenReq.RefreshToken = c.Cookies("refreshToken_cookie");
	res, err := ac.as.RefreshToken(c.Context(), refreshTokenReq)

	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
