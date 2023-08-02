package handler

import (
	"github.com/labstack/echo/v4"
	"server/model"
)

type reqCreateGenre struct {
	Title string `json:"title"`
	Color string `json:"color"`
}

type reqUpdateGenre struct {
	ID    int    `param:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

func (h *Handler) GetGenres(c echo.Context) error {
	genres, err := h.repo.GetGenres()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, genres)
}

func (h *Handler) GetGenre(c echo.Context) error {
	var req struct {
		ID       int    `param:"id"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	genre, err := h.repo.GetGenre(req.ID)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, genre)
}

func (h *Handler) CreateGenre(c echo.Context) error {
	var req reqCreateGenre
	if err := c.Bind(&req); err != nil {
		return err
	}
	genre, err := h.repo.CreateGenre(&model.Genre{
		Title: req.Title,
		Color: req.Color,
	})
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, genre)
}

func (h *Handler) UpdateGenre(c echo.Context) error {
	var req reqUpdateGenre
	if err := c.Bind(&req); err != nil {
		return err
	}

	genre, err := h.repo.GetGenre(req.ID)
	if err != nil {
		return c.JSON(500, err)
	}

	genre.Title = req.Title
	genre.Color = req.Color

	genre, err = h.repo.UpdateGenre(genre)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, genre)
}
