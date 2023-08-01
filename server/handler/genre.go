package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetGenres(c echo.Context) error {
	genres, err := h.repo.GetGenres()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, genres)
}

func (h *Handler) GetGenre(c echo.Context) error {
	var req struct {
		ID int `json:"id"`
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
