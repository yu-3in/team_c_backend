package handler

import (
	"log"
	"net/http"
	"server/model"
	"server/util"

	"github.com/labstack/echo/v4"
)

type reqCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type reqUpdateUser struct {
	ID    int    `param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type reqLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.repo.GetUsers()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, users)
}

func (h *Handler) GetUser(c echo.Context) error {
	var req struct {
		ID int `param:"id"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	user, err := h.repo.GetUser(req.ID)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, user)
}

func (h *Handler) CreateUser(c echo.Context) error {
	var req reqCreateUser
	if err := c.Bind(&req); err != nil {
		return err
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return c.JSON(500, err)
	}

	user, err := h.repo.CreateUser(&model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return c.JSON(500, err)
	}
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, echo.Map{
		"token": token,
	})
}

func (h *Handler) UpdateUser(c echo.Context) error {
	var req reqUpdateUser
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.repo.GetUser(req.ID)
	if err != nil {
		return c.JSON(500, err)
	}

	user.Name = req.Name
	user.Email = req.Email
	res, err := h.repo.UpdateUser(user)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, res)
}

func (h *Handler) Login(c echo.Context) error {
	var req reqLoginUser
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.repo.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}

	log.Println(user)

	if err := util.VerifyPassword(req.Password, user.Password); err != nil {
		return err
	}

	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (h *Handler) Logout(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
