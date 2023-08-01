package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTickets(c echo.Context) error {
	tickets, err := h.repo.GetTickets()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, tickets)
}

func (h *Handler) GetTicket(c echo.Context) error {
	var req struct {
		ID int `json:"id"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	ticket, err := h.repo.GetTicket(req.ID)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, ticket)
}
