package handler

import (
	"server/model"

	"github.com/labstack/echo/v4"
)

type reqCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type reqUpdateUser struct {
	ID      int   `param:"id"`
	Name     string `json:"name"`
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
		ID int `json:"id"`
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
	user, err := h.repo.CreateUser(&model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, user)
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