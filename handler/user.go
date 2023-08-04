package handler

import (
	"log"
	"math/rand"
	"time"
	"net/http"
	"server/handler/request"
	"server/model"
	"server/util"

	"github.com/labstack/echo/v4"
)

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

func (h *Handler) GetME(c echo.Context) error {
	userID := c.Get("userID").(int)

	user, err := h.repo.GetUser(userID)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, user)
}

func (h *Handler) CreateUser(c echo.Context) error {
	var req request.ReqCreateUser
	if err := c.Bind(&req); err != nil {
		return err
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return c.JSON(500, err)
	}

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rand.Intn(5) + 1

	colorCode := ""
	switch randomNumber {
	case 1:
		// スカーレット scarlet
		colorCode = "#ea5532"
	case 2:
		// イエロー yellow
		colorCode = "#ffdc00"
	case 3:
		// グリーン green
		colorCode = "#00a960"
	case 4:
		// ブルー blue
		colorCode = "#0075c2"
	case 5:
		// ブラック black
		colorCode = "#000000"
	}

	user, err := h.repo.CreateUser(&model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		IconColor: colorCode,
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

func (h *Handler) UpdateME(c echo.Context) error {
	userID := c.Get("userID").(int)
	var req request.ReqUpdateUser
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.repo.GetUser(userID)
	if err != nil {
		return c.JSON(500, err)
	}

	user.Name = req.Name
	user.Email = req.Email
	user.DepartmentName = req.DepartmentName
	user.ProductName = req.ProductName
	res, err := h.repo.UpdateUser(user)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, res)
}

func (h *Handler) DeleteME(c echo.Context) error {
	userID := c.Get("userID").(int)
	h.repo.DeleteMe(userID)
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) CreateUserGenre(c echo.Context) error {
	userID := c.Get("userID").(int)
	var req request.ReqUpdateUserGenre
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.repo.GetUser(userID)
	if err != nil {
		return c.JSON(500, err)
	}

	user.Genres = []model.Genre{}
	for _, genreID := range req.GenreID {
		genre, err := h.repo.GetGenre(genreID)

		if err != nil {
			return c.JSON(500, err)
		}
		user.Genres = append(user.Genres, *genre)
	}

	res, err := h.repo.CreateUserGenre(user)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, res)
}

func (h *Handler) Login(c echo.Context) error {
	var req request.ReqLoginUser
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
